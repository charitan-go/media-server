package model

// type Media struct {
// 	id         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
// 	ReadableId uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"readable_id"`
//
// 	Email string `gorm:"type:varchar(255);not null" json:"name"`
// 	// HashedPassword    string       `gorm:"type:varchar(255);not null" json:"hashed_password"`
// 	// Role              dto.RoleEnum `gorm:"type:varchar(20);not null" json:"role"`
// 	// ProfileReadableId uuid.UUID    `gorm:"type:uuid;not null" json:"profile_readable_id"`
//
// 	createdAt time.Time
// 	updatedAt time.Time
// }

// func (a *Media) BeforeCreate(db *gorm.DB) (err error) {
// 	if a.id == uuid.Nil {
// 		a.id = uuid.New()
// 	}
//
// 	if a.ReadableId == uuid.Nil {
// 		a.ReadableId = uuid.New()
// 	}
//
// 	return nil
// }
//
// func NewDonorMedia(req *dto.RegisterDonorRequestDto, hashedPassword string, role dto.RoleEnum, profileReableId uuid.UUID) *Media {
// 	return &Media{
// 		Email:             req.Email,
// 		HashedPassword:    hashedPassword,
// 		Role:              role,
// 		ProfileReadableId: profileReableId,
// 	}
// }
//
// func NewCharityMedia(req *dto.RegisterCharityRequestDto, hashedPassword string, role dto.RoleEnum, profileReableId uuid.UUID) *Media {
// 	return &Media{
// 		Email:             req.Email,
// 		HashedPassword:    hashedPassword,
// 		Role:              role,
// 		ProfileReadableId: profileReableId,
// 	}
// }
