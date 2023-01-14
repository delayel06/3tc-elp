module Main exposing (..)

import Browser
import Html exposing(..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Time

-- MAIN

main = Browser.element { init = init , update = update , subscriptions = subscriptions , view = view }


-- MODEL

type State = Success String | Failure | Loading

type alias Model = 
    {
        wordToGuess : String
        , wordSubmit : String
        , score : Int
        , timer : Int
        , httpState : State
    }

init : () -> ( Model , Cmd Msg )
init _ = 
    ( Model "hello world" "" 0 60 Loading
    , Http.get {
          url = "https://www.palabrasaleatorias.com/mots-aleatoires.php?fs=1&fs2=0&Submit=Nouveau+mot)"
        , expect = Http.expectString GotText
        }
    ) 


-- UPDATE

type Msg
    = Change String | Submit | Pass | GotText (Result Http.Error String) | Tick Time.Posix

update : Msg -> Model -> ( Model , Cmd Msg )
update msg model = 
    case msg of
        Change word -> ( { model | wordSubmit = word } , Cmd.none )

        Submit -> if model.wordSubmit == model.wordToGuess then ( { model | score = model.score + 1 , wordSubmit = "" } , Cmd.none ) else ( { model | score = model.score - 1 , wordSubmit = "" } , Cmd.none )

        Pass -> ( { model | wordSubmit = "" } , Cmd.none )

        Tick _ ->  ({ model | timer = model.timer - 1 } , Cmd.none) 
 
        GotText result -> case result of
            Ok fullText ->
                ({ model | httpState = Success fullText } , Cmd.none)

            Err _ ->
                ({ model | httpState = Failure } , Cmd.none)


-- SUBSCRIPTIONS
subscriptions : Model -> Sub Msg
subscriptions model =
    if model.timer > 0 then Time.every 1000 Tick else Sub.none

-- VIEW

view : Model -> Html Msg
view model = div [] [
      h1 [style "color" "blue" ,  style "margin-left" "50px" , style "font-family" "verdana" , style "font-size" "300%", style "width" "100%", style "align" "center"] [ text "Elmphabetic"]
    , div [ style "font-size" "150%" , style "margin" "8px 20px" ] [ div [style "float" "left", style "margin-left" "400px" , style "padding" "8px 20px"] [text ("Score : " ++ (String.fromInt model.score))] 
             , div [style "float" "left",style "padding" "8px 20px"] [text ("TimeLeft : " ++ (String.fromInt model.timer))] ]
    , div [style "width" "20%" , style "height" "50px"  , style "margin-left" "20px" , style "margin-top" "50px"] [ 
        text """
            Lorem ipsum dolor sit amet. Sit laborum quis ut voluptatem voluptas est nobis velit. Sit possimus nobis non harum natus et ipsa assumenda id voluptas 
            provident in reiciendis autem est odit tempora. Qui reprehenderit obcaecati sed expedita officiis vel delectus voluptas sit voluptatem velit At expedita 
            consequuntur qui ullam modi qui dolores dicta. Ut voluptas minus est iure quaerat qui excepturi nisi vel iste earum vel nisi officiis et fuga rerum sit
            amet fuga. Eum aliquid dolorum aut optio veritatis ad galisum veniam a fugiat dolores aut enim saepe non dolorem eaque. Et vero necessitatibus sit quisquam
            molestiae rem autem fuga et nihil quae rem fugiat internos a assumenda voluptatem. In cupiditate velit nam illum quam eum quas reprehenderit qui soluta
            aperiam et sint tempora. Qui perspiciatis voluptate sit temporibus eaque id neque neque et recusandae quidem non nihil quaerat qui distinctio ipsa qui
            dolor voluptatum! 33 sint corrupti sit suscipit praesentium ut quia cumque.
            """
        ], div[style "margin-left" "400px"] [div [] [input [ style "font-size" "20px",style "margin-left" "50px",style "margin-top" "180px", placeholder "Type your guess here !", value model.wordSubmit, onInput Change] [] ],div [] [button [ onClick Submit, style "padding" "15px 32px", style "margin-left" "50px" ] [ text "Submit"]
                , button [onClick Pass,style "margin-top" "30px", style "padding" "15px 32px", style "margin-left" "20px" ] [ text "Pass" ] ] ]]

