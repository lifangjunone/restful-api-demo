package host

import "time"

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

// Host model 定义
type Host struct {
	// 资源公共属性部分
	*Resource
	// 资源独有属性部分
	*Describe
}

func (h *Host) Validate() error {
	return validate.Struct(h)
}

func (h *Host) InjectDefault() {
	if h.CreateAt == 0 {
		h.CreateAt = time.Now().UnixMilli()
	}
}

// Vendor 云服务商
type Vendor int

const (
	// PRIVATE_IDC  枚举的默认值
	PRIVATE_IDC Vendor = iota
	// ALIYUN 阿里云
	ALIYUN
	// TXYUN 腾讯云
	TXYUN
)

// Resource 公共属性资源
type Resource struct {
	Id          string            `json:"id" validate:"required"`     // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`                     // 厂商
	Region      string            `json:"region" validate:"required"` // 地域
	CreateAt    int64             `json:"create_at"`                  // 创建时间
	ExpireAt    int64             `json:"expire_at"`                  // 创建时间
	Type        string            `json:"type" validate:"required"`   // 规格
	Name        string            `json:"name" validate:"required"`   // 名称
	Description string            `json:"description"`                // 描述
	Status      string            `json:"status"`                     // 服务商中的状态
	Tags        map[string]string `json:"tags"`                       // 标签
	UpdateAt    int64             `json:"update_at"`                  // 更新时间
	SyncAt      int64             `json:"sync_at"`                    // 同步时间
	Account     string            `json:"account"`                    // 资源所属账号
	PublicIP    string            `json:"public_ip"`                  //公网IP
	PrivateIP   string            `json:"private_ip"`                 // 内网IP
	PayType     string            `json:"pay_type"`                   // 实例支付方式
}

// Describe 独有属性资源
type Describe struct {
	CPU          int    `json:"cpu" validate:"required"`    // CPU 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU 数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU 类型
	OSType       string `json:"os_type"`                    // 操作系统类型
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
}
