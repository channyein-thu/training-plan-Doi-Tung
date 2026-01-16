package model

type CourseCategory string
type CourseType string

const (
	// Category
	CategoryEnvironment      CourseCategory = "สนับสนุนนโยบายสิ่งแวดล้อม"
	CategorySafety           CourseCategory = "ความปลอดภัยและอาชีวอนามัย"
	CategorySalesService     CourseCategory = "งานขายและงานบริการ"
	CategorySoftware         CourseCategory = "การใช้งาน Software"
	CategoryPresentation     CourseCategory = "การนำเสนอ"
	CategoryLeadership       CourseCategory = "Leadership Development"
	CategoryMachine          CourseCategory = "การใช้งานเครื่องจักรและซ่อมบำรุง"
	CategoryThinking         CourseCategory = "กระบวนการคิด วิเคราะห์"
	CategoryWorkProcess      CourseCategory = "พัฒนาทักษะกระบวนการทำงาน"
	CategoryProcurement      CourseCategory = "การจัดซื้อจัดจ้าง"
	CategoryCommunication    CourseCategory = "การสื่อสาร"
	CategorySeminar          CourseCategory = "โครงการสัมมนาอื่นๆ"
	CategoryManagement       CourseCategory = "พัฒนาขีดความสามารถระดับบริหาร"
	CategoryFinance          CourseCategory = "การเงินและการบัญชี"
)

const (
	// Type
	TypeInHouse      CourseType = "In-house"
	TypePublic       CourseType = "Public"
	TypeOJT          CourseType = "OJT"
	TypeSelfLearning CourseType = "Self-learning"
	TypeOnline       CourseType = "Online/Virtual"
)
