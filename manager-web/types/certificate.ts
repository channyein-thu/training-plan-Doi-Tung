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

export type CertificateType =
  | "In-house"
  | "Public"
  | "OJT"
  | "Self-learning"
  | "Online/Virtual";

// main type
export type Certificate = {
  id: number;
  userId: number;
  trainingName: string;
  category: CertificateCategory;
  type: CertificateType;
  image: string;
  description: string;
  createdAt: Date;
  updatedAt: Date;
};
