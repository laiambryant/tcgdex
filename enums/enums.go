package enums

type Language string

const (
LanguageEn    Language = "en"
LanguageFr    Language = "fr"
LanguageEs    Language = "es"
LanguageEsMx  Language = "es-mx"
LanguageIt    Language = "it"
LanguagePtBr  Language = "pt-br"
LanguagePtPt  Language = "pt-pt"
LanguageDe    Language = "de"
LanguageNl    Language = "nl"
LanguagePl    Language = "pl"
LanguageRu    Language = "ru"
LanguageJa    Language = "ja"
LanguageKo    Language = "ko"
LanguageZhTw  Language = "zh-tw"
LanguageId    Language = "id"
LanguageTh    Language = "th"
LanguageZhCn  Language = "zh-cn"
)

type Extension string

const (
ExtensionPng  Extension = "png"
ExtensionJpg  Extension = "jpg"
ExtensionWebp Extension = "webp"
)

type Quality string

const (
QualityLow  Quality = "low"
QualityHigh Quality = "high"
)
