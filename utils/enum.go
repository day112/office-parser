package utils

//基本类型
type BasicType string

const (
	BT_XZT   = "选择题"
	BT_DANXT = "单选题"
	BT_DUOXT = "多选题"
	BT_TKT   = "填空"
	BT_JD    = "解答"
	BT_JDT   = "解答题"
	BT_PDT   = "判断题"
	BT_ZWT   = "作文题"
	BT_TZT   = "题组题"
)

func (b BasicType) Val() string {
	switch b {
	case BT_DANXT:
		return "DANXT"
	case BT_DUOXT:
		return "DUOXT"
	case BT_TKT:
		return "TKT"
	case BT_PDT:
		return "PDT"
	case BT_JD:
		return "JD"
	case BT_JDT:
		return "JDT"
	case BT_ZWT:
		return "ZWT"
	case BT_TZT:
		return "TZT"
	case BT_XZT:
		return "XZT"
	default:
		return ""
	}
}

//应用场景
type ResUsage int

const (
	RU_ILA = iota + 1
	RU_LAS
	RU_CA
	RU_PKG
	RU_HW
)

func (r ResUsage) Val() string {
	switch r {
	case RU_ILA:
		return "ILA"
	case RU_LAS:
		return "LAS"
	case RU_CA:
		return "CA"
	case RU_PKG:
		return "PKG"
	case RU_HW:
		return "HW"
	default:
		return ""
	}
}

//structuring string
type StructuringString int

const (
	SS_DANT = 1
	//一型题组题：主题和子题使用同一级编号，1,2,3,4
	SS_TZTA = 8
	//二型题组题：主题和子题使用上下级编号，主题1，子题（1）
	SS_TZTB = 9
)

func (s StructuringString) Val() string {
	switch s {
	case SS_DANT:
		return "DANT"
	case SS_TZTA:
		return "TZTA"
	case SS_TZTB:
		return "TZTB"
	default:
		return ""
	}
}

//试题 label_string
type QuestionLabelString int

const (
	QLS_TBLXT = iota + 1
	QLS_DYT
	QLS_MNT
	QLS_ZHENT
	QLS_XBT
	QLS_YZT
	QLS_ZXT
	QLS_ZHT
)

//    TBLXT = 1  # 随堂练习题 -> 同步练习题
//    DYT = 2  # 单元测试 -> 单元测试题
//    MNT = 3  # 模拟题 -> 模拟题
//    ZHENT = 4  # 真题 -> 真题
//    XBT = 5  # 校本题 -> 校本题
//    YZT = 6  # 压轴题 -> 压轴题
//    ZXT = 7  # 专项题
//    ZHT = 8  # 综合题
func (q QuestionLabelString) Val() string {
	switch q {
	case QLS_TBLXT:
		return "TBLXT"
	case QLS_DYT:
		return "DYT"
	case QLS_MNT:
		return "MNT"
	case QLS_ZHENT:
		return "ZHENT"
	case QLS_XBT:
		return "XBT"
	case QLS_YZT:
		return "YZT"
	case QLS_ZXT:
		return "ZXT"
	case QLS_ZHT:
		return "ZHT"
	default:
		return ""
	}
}
