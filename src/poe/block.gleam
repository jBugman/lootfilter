import gleam/list
import gleam/string

import poe/actions.{type Action}
import poe/actions/label/font_size.{FontSize}
import poe/conditions.{type Condition}

pub type Block {
  Show(conditions: List(Condition), actions: List(Action))
  Hide(List(Condition))
}

const tab = "  "

const disable_drop_sound = "DisableDropSound"

pub fn to_string(block: Block) -> String {
  case block {
    Show(cond, act) -> {
      let has_sound_alert =
        list.any(act, fn(action) {
          case action {
            actions.Alert(_, _) -> True
            actions.PositionalAlert(_, _) -> True
            actions.CustomAlert(_) -> True
            _ -> False
          }
        })

      let act =
        list.append(act, [
          case has_sound_alert {
            True -> actions.DisableSoundIfAlert
            False -> actions.EnableSound
          },
        ])

      [
        ["Show"],
        cond |> indent_with(conditions.to_string),
        act |> indent_with(actions.to_string),
      ]
      |> list.flatten
      |> string.join("\n")
    }

    Hide(cond) ->
      [
        ["Hide"],
        cond |> indent_with(conditions.to_string),
        [tab <> FontSize(20) |> font_size.to_string, tab <> disable_drop_sound],
      ]
      |> list.flatten
      |> string.join("\n")
  }
}

fn indent_with(lines: List(s), f: fn(s) -> String) -> List(String) {
  list.map(lines, fn(line) { tab <> f(line) })
}
