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
import Dict exposing (Dict)


-- MAIN

main = Browser.element { init = init , update = update , subscriptions = subscriptions , view = view }


-- MODEL

type State = Success String | Failure | Loading

type alias Definition = 
    {
      def : String
    }

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
        , def : Definition
    }

init : () -> ( Model , Cmd Msg )
init _ = 
    ( Model "" "" [] 0 60 Loading Loading False (Definition "")
    , Cmd.batch [
            Http.get {
                url = "https://elm-lang.org/assets/public-opinion.txt"
                , expect = Http.expectString GotText
            }
        ]
    ) 


-- UPDATE

type Msg
    = Change String | Submit | Pass | GotText (Result Http.Error String) |GotJson (Result Http.Error Definition) | Tick Time.Posix | KeyDown RawKey | Focus | NotFocus | NewWord Int

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
                ({ model | httpState = Failure } , Cmd.none)

        Focus -> ( { model | focus = True } , Cmd.none )

        NotFocus -> ( { model | focus = False } , Cmd.none )

        NewWord id -> case (getElementAtIndex model.wordsTable id) of
                                Nothing -> (model, Cmd.none)
                                Just x -> ( { model | wordToGuess = x } , Http.get {url = ("https://api.dictionaryapi.dev/api/v2/entries/en/cat")  , expect = Http.expectJson GotJson defDecoder} )

        GotJson result -> case result of
                            Ok def -> ({ model | jsonState = Success def.def } , Cmd.none)

                            Err _ -> ({ model | jsonState = Failure } , Cmd.none)
           


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

defDecoder : Decoder Definition
defDecoder = 
    Json.Decode.map Definition (at ["0","meanings","0","definitions","0","definition"] string)



-- SUBSCRIPTIONS
subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.batch
        [ if model.timer > 0 then Time.every 1000 Tick else Sub.none, Keyboard.downs KeyDown ]


-- VIEW

view : Model -> Html Msg
view model = case  model.httpState of
    Success def -> showView model (case model.jsonState of
        Success a -> a
        Loading -> "Loading...."
        Failure -> "Can't find json")

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

isGameEnded : Model -> Bool
isGameEnded model = model.timer == 0  
