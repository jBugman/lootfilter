import block/base_types
import block/class
import block/defences
import block/quality
import block/sockets

pub type Condition {
  BaseType(base_type: base_types.BaseType)
  Class(class: class.Class)
  Defence(defence: defences.Defence)
  Quality(quality: quality.Quality)
  Sockets(sockets: sockets.Sockets)
}

pub fn to_string(cond: Condition) -> String {
  case cond {
    BaseType(x) -> base_types.to_string(x)
    Class(x) -> class.to_string(x)
    Defence(x) -> defences.to_string(x)
    Quality(x) -> quality.to_string(x)
    Sockets(x) -> sockets.to_string(x)
  }
}
