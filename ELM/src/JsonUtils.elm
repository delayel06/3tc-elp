module JsonUtils exposing (..)

import Json.Decode exposing (..)

type alias Description = List Welcome

type alias Welcome =
    { word : String
    , meanings : List Meaning
    }

type alias Meaning =
    { partOfSpeech : String
    , definitions : List Definition
    }

type alias Definition =
    { definition : String
    }

descriptionDecoder : Decoder (List Welcome)
descriptionDecoder = list welcomeDecoder 

welcomeDecoder : Decoder Welcome
welcomeDecoder = 
    map2 Welcome 
        (field "word" string)
        (field "meanings" <| list meaningDecoder)

meaningDecoder : Decoder Meaning
meaningDecoder = 
    map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" <| list definitionDecoder)

definitionDecoder : Decoder Definition
definitionDecoder = 
    map Definition 
        (field "definition" string)
        