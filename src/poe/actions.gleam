import poe/actions/beam
import poe/actions/map_icon/map_icon

pub type Action {
  MapIcon(map_icon: map_icon.MapIcon)
  Beam(beam: beam.Beam)
}

pub fn to_string(action: Action) -> String {
  case action {
    MapIcon(x) -> map_icon.to_string(x)
    Beam(x) -> beam.to_string(x)
  }
}
