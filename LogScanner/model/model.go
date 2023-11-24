package model

type Hjm3u8LogErrors struct {
	ID       uint   `gorm:"id"`
	Message  string `gorm:"message"`
	FailedAt string `gorm:"failed_at"`
	Hash     string `gorm:"hash"`
}

type ChatServerLogErrors struct {
	ID       uint   `gorm:"id"`
	Message  string `gorm:"message"`
	FailedAt string `gorm:"failed_at"`
	Hash     string `gorm:"hash"`
}

type HjAppServerErrors struct {
	ID       uint   `gorm:"id"`
	Message  string `gorm:"message"`
	FailedAt string `gorm:"failed_at"`
	Hash     string `gorm:"hash"`
}

type HjApiErrors struct {
	ID       uint   `gorm:"id"`
	Message  string `gorm:"message"`
	FailedAt string `gorm:"failed_at"`
	Hash     string `gorm:"hash"`
}

type HjAdminErrors struct {
	ID       uint   `gorm:"id"`
	Message  string `gorm:"message"`
	FailedAt string `gorm:"failed_at"`
	Hash     string `gorm:"hash"`
}

type QueueLogErrors struct {
	ID       uint   `gorm:"id"`
	Message  string `gorm:"message"`
	FailedAt string `gorm:"failed_at"`
	Hash     string `gorm:"hash"`
}
