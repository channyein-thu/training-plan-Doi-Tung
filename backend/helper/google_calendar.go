package helper

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// GOOGLE CALENDAR CLIENT
func NewGoogleCalendarService(ctx context.Context) *calendar.Service {
	credPath := viper.GetString("GOOGLE_SERVICE_ACCOUNT_FILE")

	if credPath == "" {
		log.Fatal("GOOGLE_SERVICE_ACCOUNT_FILE is not set")
	}

	credData, err := os.ReadFile(credPath)
	if err != nil {
		log.Fatal("Unable to read service account file:", err)
	}

	config, err := google.JWTConfigFromJSON(
		credData,
		calendar.CalendarEventsScope,
	)
	if err != nil {
		log.Fatal("Unable to parse service account JSON:", err)
	}

	client := config.Client(ctx)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatal("Unable to create calendar service:", err)
	}

	return srv
}

// TIMEZONE
func LoadLocation() *time.Location {
	tz := viper.GetString("TIMEZONE")
	log.Println(tz)
	if tz == "" {
		tz = "UTC"
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal("Invalid TIMEZONE:", tz)
	}

	return loc
}

// CREATE CALENDAR EVENT
func CreateTrainingPlanCalendarEvent(
	ctx context.Context,
	srv *calendar.Service,
	title string,
	description string,
	start time.Time,
	durationHours int,
) (string, error) {

	calendarID := viper.GetString("GOOGLE_CALENDAR_ID")
	if calendarID == "" {
		calendarID = "primary"
	}

	end := start.Add(time.Duration(durationHours) * time.Hour)

	event := &calendar.Event{
		Summary:     title,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: start.Format(time.RFC3339),
			TimeZone: viper.GetString("TIMEZONE"),
		},
		End: &calendar.EventDateTime{
			DateTime: end.Format(time.RFC3339),
			TimeZone: viper.GetString("TIMEZONE"),
		},
	}

	createdEvent, err := srv.Events.
		Insert(calendarID, event).
		Context(ctx).
		Do()

	if err != nil {
		return "", err
	}

	return createdEvent.Id, nil
}

// UPDATE CALENDAR EVENT
func UpdateTrainingPlanCalendarEvent(
	ctx context.Context,
	srv *calendar.Service,
	eventID string,
	title string,
	description string,
	start time.Time,
	durationHours int,
) error {

	calendarID := viper.GetString("GOOGLE_CALENDAR_ID")
	if calendarID == "" {
		calendarID = "primary"
	}

	end := start.Add(time.Duration(durationHours) * time.Hour)

	event := &calendar.Event{
		Summary:     title,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: start.Format(time.RFC3339),
			TimeZone: viper.GetString("TIMEZONE"),
		},
		End: &calendar.EventDateTime{
			DateTime: end.Format(time.RFC3339),
			TimeZone: viper.GetString("TIMEZONE"),
		},
	}

	_, err := srv.Events.
		Update(calendarID, eventID, event).
		Context(ctx).
		Do()

	return err
}

// DELETE CALENDAR EVENT
func DeleteTrainingPlanCalendarEvent(
	ctx context.Context,
	srv *calendar.Service,
	eventID string,
) error {

	calendarID := viper.GetString("GOOGLE_CALENDAR_ID")
	if calendarID == "" {
		calendarID = "primary"
	}

	return srv.Events.
		Delete(calendarID, eventID).
		Context(ctx).
		Do()
}
