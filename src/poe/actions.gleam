import gleam/int
import gleam/string

import poe/actions/beam
import poe/actions/label/border_color
import poe/actions/map_icon/map_icon
import util.{quoted}

pub type Action {
  MapIcon(map_icon: map_icon.MapIcon)
  Beam(beam: beam.Beam)

  BorderColor(color: border_color.BorderColor)

  Alert(id: Int, volume: Int)
  PositionalAlert(id: Int, volume: Int)
  CustomAlert(path: String)

  EnableSound
  DisableSoundIfAlert
}

pub fn to_string(action: Action) -> String {
  case action {
    MapIcon(x) -> map_icon.to_string(x)
    Beam(x) -> beam.to_string(x)

    BorderColor(x) -> border_color.to_string(x)

    Alert(id, volume) ->
      ["PlayAlertSound", int.to_string(id), int.to_string(volume)]
      |> string.join(" ")

    PositionalAlert(id, volume) ->
      ["PlayAlertSoundPositional", int.to_string(id), int.to_string(volume)]
      |> string.join(" ")

    CustomAlert(path) -> ["PlayCustomSound", quoted(path)] |> string.join(" ")

    EnableSound -> "EnableDropSound"
    DisableSoundIfAlert -> "DisableDropSoundIfAlertSound"
  }
}
