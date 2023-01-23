
package models

import "gorm.io/gorm"

type HeroSectionDescription struct {
    gorm.Model  // adds ID, created_at etc.
	Description string `json:"Description"`
}