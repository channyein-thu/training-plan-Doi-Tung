package helper

import (
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

// MaxCertificateFileSize caps how large an uploaded certificate file can be.
// Keep this at or below the fiber.Config.BodyLimit set in main.go.
const MaxCertificateFileSize = 5 * 1024 * 1024 // 5MB

// allowedCertificateTypes maps a lower-cased file extension to the MIME type
// we expect a genuine file of that type to sniff as. Only these extensions
// are accepted for certificate uploads.
var allowedCertificateTypes = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".webp": "image/webp",
	".pdf":  "application/pdf",
}

// ValidateUploadedFile checks an uploaded file against a size limit and an
// extension whitelist, then reads its actual leading bytes to confirm the
// content really is what the extension claims (rather than trusting the
// client-supplied filename or Content-Type header, both of which are easy
// to spoof).
//
// On success it returns the open file positioned back at the start (ready
// to be streamed to storage) along with the verified extension and MIME
// type. Callers are responsible for closing the returned file.
func ValidateUploadedFile(
	fileHeader *multipart.FileHeader,
	allowed map[string]string,
	maxSize int64,
) (file multipart.File, ext string, mimeType string, err error) {

	if fileHeader.Size <= 0 {
		return nil, "", "", BadRequest("Uploaded file is empty")
	}
	if fileHeader.Size > maxSize {
		return nil, "", "", BadRequest("File is too large (max 5MB)")
	}

	ext = strings.ToLower(filepath.Ext(fileHeader.Filename))
	expectedMime, ok := allowed[ext]
	if !ok {
		return nil, "", "", BadRequest("Unsupported file type. Allowed: JPG, PNG, WEBP, PDF")
	}

	f, err := fileHeader.Open()
	if err != nil {
		return nil, "", "", BadRequest("Failed to open uploaded file")
	}

	sniff := make([]byte, 512)
	n, readErr := io.ReadFull(f, sniff)
	if readErr != nil && readErr != io.ErrUnexpectedEOF && readErr != io.EOF {
		f.Close()
		return nil, "", "", BadRequest("Failed to read uploaded file")
	}

	detected := http.DetectContentType(sniff[:n])
	if !contentTypeMatches(detected, expectedMime) {
		f.Close()
		return nil, "", "", BadRequest("File content does not match its extension")
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		f.Close()
		return nil, "", "", Internal("Failed to process uploaded file")
	}

	return f, ext, expectedMime, nil
}

// ValidateCertificateFile is a convenience wrapper around ValidateUploadedFile
// using the certificate-upload whitelist and size cap.
func ValidateCertificateFile(fileHeader *multipart.FileHeader) (multipart.File, string, string, error) {
	return ValidateUploadedFile(fileHeader, allowedCertificateTypes, MaxCertificateFileSize)
}

func contentTypeMatches(detected, expected string) bool {
	// http.DetectContentType can append parameters (e.g. "text/plain; charset=utf-8")
	// for text-like content; none of our allowed binary types do this, but strip
	// defensively so a stray parameter never causes a false negative.
	if idx := strings.Index(detected, ";"); idx != -1 {
		detected = detected[:idx]
	}
	return strings.EqualFold(strings.TrimSpace(detected), expected)
}
