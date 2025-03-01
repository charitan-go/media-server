package dto

type MediaTypeEnum string

const (
	MediaTypeImage MediaTypeEnum = "IMAGE"
	MediaTypeVideo MediaTypeEnum = "VIDEO"
)

// Values provides list valid values for Enum.
func (MediaTypeEnum) Values() (kinds []string) {
	for _, s := range []MediaTypeEnum{
		MediaTypeImage,
		MediaTypeVideo,
	} {
		kinds = append(kinds, string(s))
	}
	return
}
