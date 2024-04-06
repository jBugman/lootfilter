import gleam/erlang
import gleam/erlang/atom.{type Atom}
import gleam/erlang/charlist.{type Charlist}
import gleam/result

// {ok, [[Home]]} = init:get_argument(home)
@external(erlang, "init", "get_argument")
fn init_get_argument(flag: Atom) -> Result(Charlist, a)

pub fn homedir() -> Result(String, String) {
  atom.create_from_string("home")
  |> init_get_argument
  |> result.map(charlist.to_string)
  |> result.map_error(erlang.format)
}
