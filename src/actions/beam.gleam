import actions/color.{type Color}

pub type Beam {
  Beam(Color)
}

pub fn to_string(beam: Beam) -> String {
  let Beam(c) = beam
  "PlayEffect " <> color.to_string(c)
}
