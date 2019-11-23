/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 21:47:55
 * @LastEditors: your name
 */
package db

import (
	"time"
)

// Version 存储系统版本信息 PASS
type Version struct {
	ID          int       `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	VersionNo   string    `gorm:"type:varchar(8);column:version_no"`
	ReleaseTime time.Time `gorm:"column:release_time"`
	ReleaseNote string    `gorm:"column:release_note"`
}

// TableName 指定Version表名
func (Version) TableName() string {
	return "version"
}

// Role 角色
type Role struct {
	ID         uint   `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	Name       string `gorm:"type:varchar(64);unique"`
	Permission int    `gorm:"column:permission"`
	Default    bool   `gorm:"column:default;default:false;index"`
	Users      []User `gorm:"foreignkey:RoleID"`
}

// TableName 指定Role表名
func (Role) TableName() string {
	return "roles"
}

// User 用户
type User struct {
	ID           int       `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	Email        string    `gorm:"column:email;type:varchar(64);unique;not null;index"`
	Username     string    `gorm:"column:username;type:varchar(64)"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(128)"`
	RegisterTime time.Time `gorm:"column:registertime"`
	LastSeen     time.Time `gorm:"column:last_seen"`
	RoleID       uint      `gorm:"column:role_id"`
	Role         Role      `gorm:"ForeignKey:RoleID;ASSOCIATION_FOREIGNKEY:ID"`
	DataSet      []DataSet `gorm:"foreignkey:CreatorID"`
	Job          []Job     `gorm:"foreignkey:CreatorID"`
}

// TableName 指定User表名
func (User) TableName() string {
	return "user"
}

// DataSet 数据集
type DataSet struct {
	ID          int       `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	URL         string    `gorm:"column:url"`
	FileName    string    `gorm:"type:text;column:filename;unique"`
	FileHash    string    `gorm:"type:varchar(64);column:filehash"`
	FileType    string    `gorm:"type:varchar(8);column:filetype"`
	FileSize    uint      `gorm:"column:filezise"`
	FileNumer   uint      `gorm:"column:filenumber"`
	Description string    `gorm:"column:description"`
	CreatorID   uint      `gorm:"column:creator"`
	Creator     User      `gorm:"FOREIGNKEY:CreatorID;ASSOCIATION_FOREIGNKEY:ID"`
	CreatedDate time.Time `gorm:"column:created_date;default:CURRENT_TIMESTAMP"`
	Shared      bool      `gorm:"column:shared;default:false"`
}

// TableName 指定DataSet表名
func (DataSet) TableName() string {
	return "dataset"
}

// Job 任务
type Job struct {
	ID uint `gorm:"AUTO_INCREMENT;column:id;primary_key"`

	Name        string `gorm:"column:name;type:varchar(32)"`
	Namespace   string `gorm:"column:namespace;type:varchar(32)"`
	HoastName   string `gorm:"column:hostname;type:varchar(16)"`
	Container   string `gorm:"column:container;type:varchar(64)"`
	Description string `gorm:"column:description"`
	WorkENV     string `gorm:"column:work_env;type:varchar(32)"`
	DataSets    string `gorm:"column:datasets"`
	State       string `gorm:"column:state;type:varchar(8)"`
	// params = db.Column(db.JSON)
	// 这个需要更正为JSON类型
	Params      string `gorm:"params"`
	Notes       string `gorm:"column:notes"`
	Command     string `gorm:"column:command"`
	GitURL      string `gorm:"column:git_url;type:varchar(64)"`
	GitBranch   string `gorm:"column:git_branch;type:varchar(32)"`
	GitCommit   string `gorm:"column:git_commit;type:varchar(64)"`
	DockerImage string `gorm:"column:docker_image;type:varchar(32)"`
	UseGPU      bool   `gotm:"column:use_gpu"`
	// 这个需要更正为JSON
	Definition  string    `gorm:"column:definition"`
	StartTime   time.Time `gorm:"column:start_time"`
	EndTime     time.Time `gorm:"column:end_time"`
	CreatedDate time.Time `gorm:"column:create_date;default:CURRENT_TIMESTAMP"`
	// 这个需要更正为JSON
	Environment string  `gorm:"column:environment"`
	Callback    string  `gorm:"column:callback;type:varchar(64)"`
	Deletion    bool    `gorm:"column:deletion"`
	Category    string  `gorm:"column:category;type:varchar(16)"`
	CreatorID   uint    `gorm:"column:creator"`
	Creator     User    `gorm:"FOREIGNKEY:CreatorID;ASSOCIATION_FOREIGNKEY:ID"`
	Event       []Event `gorm:"foreignkey:JobID"`
}

// TableName 指定Job表名
func (Job) TableName() string {
	return "job"
}

// Event 事件
type Event struct {
	ID          uint      `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	EventName   string    `gorm:"column:event_name;type:varchar(32)"`
	Message     string    `gorm:"column:message"`
	CreatedDate time.Time `gorm:"column:create_date;default:CURRENT_TIMESTAMP"`
	JobID       uint      `gorm:"column:job"`
	Job         Job       `gorm:"ForeignKey:JobID;ASSOCIATION_FOREIGNKEY:ID"`
}

// TableName 指定Event表名
func (Event) TableName() string {
	return "event"
}

// Images Docker镜像
type Images struct {
	ID          uint   `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	Name        string `gorm:"column:name;type:varchar(128);unique"`
	Version     string `gorm:"column:version;type:varchar(8)"`
	Description string `gorm:"column:description"`
	// 这个需要改为JSON
	Include     string    `gorm:"column:include"`
	CreatedDate time.Time `gorm:"column:create_date"`
}

// TableName 指定Images表名
func (Images) TableName() string {
	return "images"
}

// Station x
type Station struct {
	ID   uint   `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	Name string `gorm:"column:name;type:varchar(32);unique"`
	// 这个需要改为JSON
	Deatil string `gorm:"column:detail"`
	Status bool   `gorm:"column:status;default:true"`
}

// TableName 指定Station表名
func (Station) TableName() string {
	return "stations"
}

// Resource x
type Resource struct {
	ID          uint      `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	Node        Station   `gorm:"ForeignKey:ID;column:node"`
	CollectTime time.Time `gorm:"column:collect_time"`
	// 这个需要改成JSON
	Detail string `gorm:"column:detail"`
}

// TableName 指定Resource表名
func (Resource) TableName() string {
	return "Resources"
}

// Dispatcher 调度器
type Dispatcher struct {
	ID uint `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	// 这个需要改成JSON
	Threshold string `gorm:"column:threshold"`
	// 这个需要改成JSON
	Job string `gorm:"column:job"`
	// 这个需要改成JSON
	SortParams   string `gorm:"column:sort_params"`
	LoopInterval uint   `gorm:"column:loop_interval"`
}

// TableName 指定Dispatcher表名
func (Dispatcher) TableName() string {
	return "dispatcher"
}

// Archive 文件
type Archive struct {
	ID          uint      `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	GlacierID   string    `gorm:"column:glacier_id;varchar(138)"`
	Checksum    string    `gorm:"column:checksum;type:varchar(138)"`
	FileName    string    `gorm:"column:filename;type:varchar(150)"`
	FileType    string    `gorm:"column:filetype;type:varchar(8)"`
	FileSize    uint      `gorm:"column:filesize"`
	Description string    `gorm:"column:description"`
	CreatorID   uint      `gorm:"column:creator"`
	Creator     User      `gorm:"ForeignKey:CreatorID;ASSOCIATION_FOREIGNKEY:ID"`
	CreatedDate time.Time `gorm:"column:create_date;default:CURRENT_TIMESTAMP"`
	Shared      bool      `gorm:"columm:shared;default:false"`
}

// TableName 指定Archive表名
func (Archive) TableName() string {
	return "archive"
}

// RPCService rpc 服务 PASS
type RPCService struct {
	ID   uint   `gorm:"AUTO_INCREMENT;column:id;primary_key"`
	Name string `gorm:"column:name;type:varchar(16),primary_key"`
	//user:pwd@host:port
	URI       string `gorm:"column:uri;type:varchar(48)"`
	Input     string `gorm:"column:input;type:varchar(8)"`
	ParamKeys string `gorm:"column:param_keys;type:varchar(32)"`
}

// TableName 指定RPCService表名
func (RPCService) TableName() string {
	return "rpc_service"
}

// StatisticsItem 统计项 PASS
type StatisticsItem struct {
	Item string `gorm:"column:item;type:varchar(16);primary_key"`
	// 这个需要换成JSON
	Detail string `gorm:"column:detail"`
}

// TableName 指定StatisticsItem的表名
func (StatisticsItem) TableName() string {
	return "statistics_item"
}
