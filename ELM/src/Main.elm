module Main exposing (main)

import Browser
import Html exposing (Html, div, text, Attribute, input)
import Html.Attributes exposing (class, placeholder)
import Html.Events exposing (onInput)
import String
import Html.Attributes exposing (..)


-- MODEL

type alias Model =
    { word : String
    , definition : String
    , guess : String
    , message : String
    }


initialModel : Model
initialModel =
    { word = "elm"
    , definition = "A functional programming language for the web"
    , guess = ""
    , message = ""
    }


-- UPDATE

type Msg
    = UpdateGuess String


update : Msg -> Model -> Model
update msg model =
    case msg of
        UpdateGuess guess ->
            { model | guess = guess }


-- VIEW

view : Model -> Html Msg
view model =
    div []
        [ div [ class "definition" ] [ text model.definition ]
        , input [ placeholder "Enter your guess here", onInput UpdateGuess ] []
        , div [ class "message" ] [ text model.message ]
        ]


-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none


-- MAIN

main =
--  Browser.element
--  { 
--      init = initialModel
--      , view = view
--      , update = update
--      , subscriptions = subscriptions
--  }

    Browser.sandbox
    {
        init = initialModel
        , view = view
        , update = update
    }
