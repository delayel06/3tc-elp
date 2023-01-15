module Utils exposing (..)

import Html exposing(..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Time
import Keyboard exposing (RawKey)
import Random
import Json.Decode exposing (..)
import JsonUtils exposing (..)


-- MODEL

type State = Success String | Failure | Loading

type alias Model = 
    {
          wordToGuess : String
        , wordSubmit : String
        , wordsTable : List String
        , score : Int
        , timer : Int
        , httpState : State
        , jsonState : State
        , focus : Bool
    }


-- Message
type Msg
    = Change String | Submit | Pass | GotText (Result Http.Error String) |GotJson (Result Http.Error Description) | Tick Time.Posix | KeyDown RawKey | Focus | NotFocus | NewWord Int


-- HELPERS

checkSubmit : Model -> ( Model , Cmd Msg )
checkSubmit model =
    if model.wordSubmit == model.wordToGuess
        then ( { model | score = model.score + 1 , wordSubmit = "" } , Random.generate NewWord (Random.int 1 1000) )
        else ( { model | score = model.score - 1 , wordSubmit = "" } , Random.generate NewWord (Random.int 1 1000) )

getElementAtIndex : List a -> Int -> Maybe a
getElementAtIndex list index =
    if index < 0 || index >= List.length list then
        Nothing
    else
        List.head (List.drop index list)

isGameEnded : Model -> Bool
isGameEnded model = model.timer == 0  


-- HTML DISPLAY

showView : Model -> String -> Html Msg
showView model definition = 
    div [] 
        [
          h1 [style "color" "blue" ,  style "margin-left" "50px" , style "font-family" "verdana" , style "font-size" "300%", style "width" "100%", style "align" "center"]
             [ text "Elmphabetic"]
        , div [ style "font-size" "150%" , style "margin" "8px 20px" ] 
            [ div [style "float" "left", style "margin-left" "500px" , style "padding" "8px 20px"]
                 [text ("Score : " ++ (String.fromInt model.score))] 
            , div [style "float" "left",style "padding" "8px 20px"]
                 [text ("TimeLeft : " ++ (String.fromInt model.timer))] ]
        , div [style "width" "25%" , style "height" "50px"  , style "margin-left" "20px" , style "margin-top" "80px", style "font-size" "18px"] 
            [ 
            text definition
            ]
        , div [ style "margin-left" "500px" ] 
            [ div [] 
                [ input [ style "font-size" "20px",style "margin-left" "50px",style "margin-top" "180px", placeholder "Type your guess here !", Html.Attributes.value model.wordSubmit, onInput Change, onFocus Focus ,onBlur NotFocus , disabled (isGameEnded model )] [] ]
            , div [] 
                [ button [ onClick Submit, style "padding" "15px 32px", style "margin-left" "50px" , disabled ( isGameEnded model )]
                    [ text "Submit"]
                , button [onClick Pass,style "margin-top" "30px", style "padding" "15px 32px", style "margin-left" "20px" , disabled (isGameEnded model )] 
                    [ text "Pass" ]
                ]
            ]
        , div [ style "font-size" "48px" , style "margin" "100px", style "color" "green"]
            [ text (if isGameEnded model then ( "Bravo !! Votre score est de " ++ String.fromInt model.score ++ ", rechargez la page pour réessayer") else "")]   
        ]