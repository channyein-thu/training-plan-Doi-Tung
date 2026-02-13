// enums
export type CertificateCategory =
  | "สนับสนุนนโยบายสิ่งแวดล้อม"
  | "ความปลอดภัยและอาชีวอนามัย"
  | "งานขายและงานบริการ"
  | "การใช้งาน Software"
  | "การนำเสนอ"
  | "Leadership Development"
  | "การใช้งานเครื่องจักรและซ่อมบำรุง"
  | "กระบวนการคิด วิเคราะห์"
  | "พัฒนาทักษะกระบวนการทำงาน"
  | "การจัดซื้อจัดจ้าง"
  | "การสื่อสาร"
  | "โครงการสัมมนาอื่นๆ"
  | "พัฒนาขีดความสามารถระดับบริหาร"
  | "การเงินและการบัญชี";

export type CertificateStatus = "Pending" | "Approved" | "Rejected";

// main type
export type Certificate = {
  id: number;
  userId: number;
  userName: string;
  employeeId: string;
  department: string;
  division: string;
  category: CertificateCategory;
  trainingId: number;
  trainingName: string;
  image: string;
  description: string;
  status: CertificateStatus;
  createdAt: string; // ISO string from backend
  updatedAt: string; // ISO string from backend
};
