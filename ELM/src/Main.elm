module Main exposing (..)

import Browser
import Html exposing(..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)

-- MAIN

main = Browser.sandbox {init = init, update = update, view = view}


-- MODEL

type alias Model = 
    {
        wordSubmit : String
        , score : Int
        , timer : Int
    }

init : Model
init = 
    Model "" 0 60


-- UPDATE

type Msg
    = Change String | IncrementScore Int | DecrementScore Int | IncrementTimer Int | DecrementTimer Int | Submit | Pass

update : Msg -> Model -> Model
update msg model = 
    case msg of
        Change word -> { model | wordSubmit = word } 

        IncrementScore val -> { model | score = model.score + 1 }

        DecrementScore val -> { model | score = model.score - 1 }

        IncrementTimer val -> { model | timer = model.timer + 1 }

        DecrementTimer val -> { model | timer = model.timer - 1 }

        Submit -> { model | wordSubmit = "" } 

        Pass -> { model | wordSubmit = "" }


-- VIEW

view : Model -> Html Msg
view model = div [] [
      header [] [ text "Elmphabetic"]
    , div [] [ text ("Score : " ++ (String.fromInt model.score))
             , text ("TimeLeft : " ++ (String.fromInt model.timer))]
    , div [] [ 
        text """
            Lorem ipsum dolor sit amet. Sit laborum quis ut voluptatem voluptas est nobis velit. Sit possimus nobis non harum natus et ipsa assumenda id voluptas 
            provident in reiciendis autem est odit tempora. Qui reprehenderit obcaecati sed expedita officiis vel delectus voluptas sit voluptatem velit At expedita 
            consequuntur qui ullam modi qui dolores dicta. Ut voluptas minus est iure quaerat qui excepturi nisi vel iste earum vel nisi officiis et fuga rerum sit
            amet fuga. Eum aliquid dolorum aut optio veritatis ad galisum veniam a fugiat dolores aut enim saepe non dolorem eaque. Et vero necessitatibus sit quisquam
            molestiae rem autem fuga et nihil quae rem fugiat internos a assumenda voluptatem. In cupiditate velit nam illum quam eum quas reprehenderit qui soluta
            aperiam et sint tempora. Qui perspiciatis voluptate sit temporibus eaque id neque neque et recusandae quidem non nihil quaerat qui distinctio ipsa qui
            dolor voluptatum! 33 sint corrupti sit suscipit praesentium ut quia cumque.
            """
        , input [ placeholder "Type your guess here !", value model.wordSubmit, onInput Change] [] ]
        , div [] [ button [ onClick Submit] [ text "Submit"]
                 , button [onClick Pass ] [ text "Pass" ] ] ]
