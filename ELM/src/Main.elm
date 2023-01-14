module Main exposing (..)

import Browser
import Html exposing(..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Time
import Keyboard exposing (RawKey)
import Random

-- MAIN

main = Browser.element { init = init , update = update , subscriptions = subscriptions , view = view }


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
        , focus : Bool
    }

init : () -> ( Model , Cmd Msg )
init _ = 
    ( Model "" "" [] 0 180 Loading False
    , Http.get {
          url = "https://elm-lang.org/assets/public-opinion.txt"
        , expect = Http.expectString GotText
        }

    ) 


-- UPDATE

type Msg
    = Change String | Submit | Pass | GotText (Result Http.Error String) | Tick Time.Posix | KeyDown RawKey | Focus | NotFocus | NewWord Int

update : Msg -> Model -> ( Model , Cmd Msg )
update msg model = 
    case msg of
        Change word -> ( { model | wordSubmit = word } , Cmd.none )

        Submit -> checkSubmit model

        KeyDown key -> if (Keyboard.anyKeyOriginal key) == Just Keyboard.Enter && model.focus == True then checkSubmit model else (model , Cmd.none)

        Pass -> ( { model | wordSubmit = "" } , Cmd.none )

        Tick _ ->  ({ model | timer = model.timer - 1 } , Cmd.none) 
 
        GotText result -> case result of
            Ok fullText ->
                ({ model | wordsTable = String.split " " fullText , httpState = Success fullText } , Cmd.none)

            Err _ ->
                ({ model | httpState = Failure } , Cmd.none)

        Focus -> ( { model | focus = True } , Cmd.none )

        NotFocus -> ( { model | focus = False } , Cmd.none )

        NewWord id -> case (getElementAtIndex model.wordsTable id) of
                                Nothing -> (model, Cmd.none)
                                Just x -> ( { model | wordToGuess = x } , Cmd.none )


checkSubmit : Model -> ( Model , Cmd Msg )
checkSubmit model =
    if model.wordSubmit == model.wordToGuess
        then ( { model | score = model.score + 1 , wordSubmit = "" } , Random.generate NewWord (Random.int 1 10) )
        else ( { model | score = model.score - 1 , wordSubmit = "" } , Random.generate NewWord (Random.int 1 10) )

getElementAtIndex : List a -> Int -> Maybe a
getElementAtIndex list index =
    if index < 0 || index >= List.length list then
        Nothing
    else
        List.head (List.drop index list)

-- SUBSCRIPTIONS
subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.batch
        [ if model.timer > 0 then Time.every 1000 Tick else Sub.none, Keyboard.downs KeyDown ]


-- VIEW

view : Model -> Html Msg
view model = case  model.httpState of
    Success def -> showView model model.wordToGuess

    Loading -> showView model "Loading..."

    Failure -> showView model "Can't access the page"

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
                [ input [ style "font-size" "20px",style "margin-left" "50px",style "margin-top" "180px", placeholder "Type your guess here !", value model.wordSubmit, onInput Change, onFocus Focus ,onBlur NotFocus] [] ]
            , div [] 
                [ button [ onClick Submit, style "padding" "15px 32px", style "margin-left" "50px" ]
                    [ text "Submit"]
                , button [onClick Pass,style "margin-top" "30px", style "padding" "15px 32px", style "margin-left" "20px" ] 
                    [ text "Pass" ]
                ]
            ]
        ]