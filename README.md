# Free Translate API

Welcome to the Free Translate API! This API allows you to perform translations between different languages without any limits. It's a free and open-source project ported and baked by
[Ismal Zikri](https://twitter.com/ZikriIsmal) .

## Table of Contents

- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [Supported Languages](#supported-languages)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Installation

To use the Free Translate API, you need to have [Golang](https://golang.org/dl/) installed on your system. You can then follow these steps to get started:

 Clone the repository:

   ```bash
   git clone https://github.com/ismalzikri/free-translate-api.git

   cd free-translate-api

   ```

## usage
Start the API server:
```bash
go run main.go
```
The API server will start running on http://localhost:8000.

You can now send translation requests to the API using HTTP POST requests. The endpoint for translation is /translate. Here's an example using curl:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"text": "Hello", "to": "id"}' http://localhost:8000/translate
```

## Supported Languages
The Free Translate API supports translation between a wide range of languages. You can check this a list of all supported languages and their language codes.
```
AFAR                = "aa"
ABKHAZIAN           = "ab"
AFRIKAANS           = "af"
AKAN                = "ak"
AMHARIC             = "am"
ARAGONESE           = "an"
ARABIC              = "ar"
ASSAMESE            = "as"
AVAR                = "av"
AYMARA              = "ay"
AZERBAIJANI         = "az"
BASHKIR             = "ba"
BELARUSIAN          = "be"
BULGARIAN           = "bg"
BIHARI              = "bh"
BISLAMA             = "bi"
BAMBARA             = "bm"
BENGALI             = "bn"
TIBETAN             = "bo"
BRETON              = "br"
BOSNIAN             = "bs"
CATALAN             = "ca"
CHECHEN             = "ce"
CHAMORRO            = "ch"
CORSICAN            = "co"
CREE                = "cr"
CZECH               = "cs"
OLD_CHURCH_SLAVONIC = "cu"
OLD_BULGARIAN       = "cu"
CHUVASH             = "cv"
WELSH               = "cy"
DANISH              = "da"
GERMAN              = "de"
DIVEHI              = "dv"
DZONGKHA            = "dz"
EWE                 = "ee"
GREEK               = "el"
ENGLISH             = "en"
ESPERANTO           = "eo"
SPANISH             = "es"
ESTONIAN            = "et"
BASQUE              = "eu"
PERSIAN             = "fa"
PEUL                = "ff"
FINNISH             = "fi"
FIJIAN              = "fj"
FAROESE             = "fo"
FRENCH              = "fr"
WEST_FRISIAN        = "fy"
IRISH               = "ga"
SCOTTISH_GAELIC     = "gd"
GALICIAN            = "gl"
GUARANI             = "gn"
GUJARATI            = "gu"
MANX                = "gv"
HAUSA               = "ha"
HEBREW              = "he"
HINDI               = "hi"
HIRI_MOTU           = "ho"
CROATIAN            = "hr"
HAITIAN             = "ht"
HUNGARIAN           = "hu"
ARMENIAN            = "hy"
HERERO              = "hz"
INTERLINGUA         = "ia"
INDONESIAN          = "id"
INTERLINGUE         = "ie"
IGBO                = "ig"
SICHUAN_YI          = "ii"
INUPIAK             = "ik"
IDO                 = "io"
ICELANDIC           = "is"
ITALIAN             = "it"
INUKTITUT           = "iu"
JAPANESE            = "ja"
JAVANESE            = "jv"
GEORGIAN            = "ka"
KONGO               = "kg"
KIKUYU              = "ki"
KUANYAMA            = "kj"
KAZAKH              = "kk"
GREENLANDIC         = "kl"
CAMBODIAN           = "km"
KANNADA             = "kn"
KOREAN              = "ko"
KANURI              = "kr"
KASHMIRI            = "ks"
KURDISH             = "ku"
KOMI                = "kv"
CORNISH             = "kw"
KIRGHIZ             = "ky"
LATIN               = "la"
LUXEMBOURGISH       = "lb"
GANDA               = "lg"
LIMBURGIAN          = "li"
LINGALA             = "ln"
LAOTIAN             = "lo"
LITHUANIAN          = "lt"
LATVIAN             = "lv"
MALAGASY            = "mg"
MARSHALLESE         = "mh"
MAORI               = "mi"
MACEDONIAN          = "mk"
MALAYALAM           = "ml"
MONGOLIAN           = "mn"
MOLDOVAN            = "mo"
MARATHI             = "mr"
MALAY               = "ms"
MALTESE             = "mt"
BURMESE             = "my"
NAURUAN             = "na"
NORTH_NDEBELE       = "nd"
NEPALI              = "ne"
NDONGA              = "ng"
DUTCH               = "nl"
NORWEGIAN_NYNORSK   = "nn"
NORWEGIAN           = "no"
SOUTH_NDEBELE       = "nr"
NAVAJO              = "nv"
CHICHEWA            = "ny"
OCCITAN             = "oc"
OJIBWA              = "oj"
OROMO               = "om"
ORIYA               = "or"
OSSETIAN            = "os"
PUNJABI             = "pa"
PALI                = "pi"
POLISH              = "pl"
PASHTO              = "ps"
PORTUGUESE          = "pt"
QUECHUA             = "qu"
RAETO_ROMANCE       = "rm"
KIRUNDI             = "rn"
ROMANIAN            = "ro"
RUSSIAN             = "ru"
RWANDI              = "rw"
SANSKRIT            = "sa"
SARDINIAN           = "sc"
SINDHI              = "sd"
SANGO               = "sg"
SERBO_CROATIAN      = "sh"
SINHALESE           = "si"
SLOVAK              = "sk"
SLOVENIAN           = "sl"
SAMOAN              = "sm"
SHONA               = "sn"
SOMALIA             = "so"
ALBANIAN            = "sq"
SERBIAN             = "sr"
SWATI               = "ss"
SOUTHERN_SOTHO      = "st"
SUNDANESE           = "su"
SWEDISH             = "sv"
SWAHILI             = "sw"
TAMIL               = "ta"
TELUGU              = "te"
TAJIK               = "tg"
THAI                = "th"
TIGRINYA            = "ti"
TURKMEN             = "tk"
TAGALOG             = "tl"
TSWANA              = "tn"
TONGA               = "to"
TURKISH             = "tr"
TSONGA              = "ts"
TATAR               = "tt"
TWI                 = "tw"
TAHITIAN            = "ty"
UYGHUR              = "ug"
UKRAINIAN           = "uk"
URDU                = "ur"
VENDA               = "ve"
VIETNAMESE          = "vi"
VOLAPÜK             = "vo"
WALLOON             = "wa"
WOLOF               = "wo"
XHOSA               = "xh"
YIDDISH             = "yi"
YORUBA              = "yo"
ZHUANG              = "za"
CHINESE             = "zh"
ZULU                = "zu"
```

## API Endpoints
`POST /translate`: 

Translates the provided text to the specified target language. The request body should be a JSON object with the following structure:
```json
{
  "text": "Hello",
  "to": "id"
}
```
The response will contain the translated text.

## Contributing
Contributions to the Free Translate API are always welcome! If you find any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License.
