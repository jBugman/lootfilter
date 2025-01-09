import gleam/int

pub type Sound {
  Mute
  Enable
  Alert(Int)
}

pub fn render(sound: Sound) -> List(String) {
  case sound {
    Mute -> ["DisableDropSound"]
    Enable -> ["EnableDropSound"]
    Alert(num) -> ["DisableDropSound", "PlayAlertSound " <> int.to_string(num)]
  }
}
