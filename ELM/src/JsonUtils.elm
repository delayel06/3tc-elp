module JsonUtils exposing (..)

import Html exposing(..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Json.Decode exposing (..)
import Utils exposing (..)


descriptionDecoder : Decoder (List Welcome)
descriptionDecoder = Json.Decode.list welcomeDecoder 

welcomeDecoder : Decoder Welcome
welcomeDecoder = 
    map2 Welcome 
        (field "word" string)
        (field "meanings" <| Json.Decode.list meaningDecoder)

meaningDecoder : Decoder Meaning
meaningDecoder = 
    map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" <| Json.Decode.list definitionDecoder)

definitionDecoder : Decoder Definition
definitionDecoder = 
    Json.Decode.map Definition 
        (field "definition" string)

displayDesc : Description -> List ( Html Msg )
displayDesc list = case list of
    [] -> []
    (x::xs) -> [li [] ([text "Meaning"] ++ [ul [] (displayMeanings x.meanings)])] ++ (displayDesc xs) 

displayMeanings : List Meaning -> List (Html Msg)
displayMeanings list = case list of
    [] -> []
    (x::xs) -> [li [] [text x.partOfSpeech]] ++ [ol [] (displayDefinitions x.definitions)] ++ (displayMeanings xs)

displayDefinitions : List Definition -> List (Html Msg)
displayDefinitions list = case list of
    [] -> []
    (x::xs) -> [li [] [text x.definition]] ++ (displayDefinitions xs)

   