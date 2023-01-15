module Main exposing (..)

import Browser
import Html exposing(..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Time
import Keyboard exposing (RawKey)
import Random
import Json.Decode exposing (..)
import Utils exposing (..)
import JsonUtils exposing (..)


-- MAIN

main = Browser.element { init = init , update = update , subscriptions = subscriptions , view = view }


-- INIT

init : () -> ( Model , Cmd Msg )
init _ = 
    ( Model "" "" [] 0 60 Loading Loading False
    , Cmd.batch [
            Http.get {
                url = "https://elm-lang.org/assets/public-opinion.txt"
                , expect = Http.expectString GotText
            }
        ]
    ) 


-- UPDATE


update : Msg -> Model -> ( Model , Cmd Msg )
update msg model = 
    case msg of
        Change word -> ( { model | wordSubmit = word } , Cmd.none )

        Submit -> checkSubmit model

        KeyDown key -> if (Keyboard.anyKeyOriginal key) == Just Keyboard.Enter && model.focus == True then checkSubmit model else (model , Cmd.none)

        Pass -> ( { model | wordSubmit = "" } , Random.generate NewWord (Random.int 1 1000) )

        Tick _ ->  ({ model | timer = model.timer - 1 } , Cmd.none) 
 
        GotText result -> case result of
            Ok fullText ->
                ({ model | wordsTable = String.split " " fullText , httpState = Success fullText } , Random.generate NewWord (Random.int 1 1000))

            Err _ ->
                ({ model | httpState = Utils.Failure } , Cmd.none)

        Focus -> ( { model | focus = True } , Cmd.none )

        NotFocus -> ( { model | focus = False } , Cmd.none )

        NewWord id -> case (getElementAtIndex model.wordsTable id) of
                                Nothing -> (model, Cmd.none)
                                Just x -> ( { model | wordToGuess = x } , Http.get {url = ("https://api.dictionaryapi.dev/api/v2/entries/en/cat")  , expect = Http.expectJson GotJson descriptionDecoder} )

        GotJson result -> case result of
                            Ok desc -> ({ model | jsonState = Success "Ok" } , Cmd.none)

                            Err _ -> ({ model | jsonState = Utils.Failure } , Cmd.none)
           


-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.batch
        [ if model.timer > 0 then Time.every 1000 Tick else Sub.none, Keyboard.downs KeyDown ]


-- VIEW

view : Model -> Html Msg
view model = case  model.httpState of
    Utils.Success def -> showView model (case model.jsonState of
        Utils.Success a -> model.wordToGuess
        Utils.Loading -> "Loading...."
        Utils.Failure -> "Can't find json")

    Utils.Loading -> showView model "Loading..."

    Utils.Failure -> showView model "Can't access the page"