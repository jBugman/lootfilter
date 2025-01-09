pub type Shape {
  Circle
  Cross
  Diamond
  Hexagon
  Kite
  Moon
  Pentagon
  Raindrop
  Square
  Star
  Triangle
  UpsideDownHouse
}

pub fn to_string(shape: Shape) -> String {
  case shape {
    Circle -> "Circle"
    Cross -> "Cross"
    Diamond -> "Diamond"
    Hexagon -> "Hexagon"
    Kite -> "Kite"
    Moon -> "Moon"
    Pentagon -> "Pentagon"
    Raindrop -> "Raindrop"
    Square -> "Square"
    Star -> "Star"
    Triangle -> "Triangle"
    UpsideDownHouse -> "UpsideDownHouse"
  }
}
