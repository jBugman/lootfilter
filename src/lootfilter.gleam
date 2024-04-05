import argv
import config
import envoy
import filepath
import gleam/io
import gleam/list
import gleam/result
import simplifile

pub fn main() {
  out_dir()
  |> io.println

  case argv.load().arguments {
    [path_in] -> build(path_in)
    _ -> io.println("usage: app [input_file]")
  }
}

const poe_dir_path = ["Documents", "My Games", "Path of Exile"]

fn out_dir() -> String {
  let homedir =
    envoy.get("HOME")
    |> result.unwrap("")

  list.fold(poe_dir_path, homedir, filepath.join)
}

fn build(src_path: String) {
  let assert Ok(data) = simplifile.read(src_path)
  config.parse(data)
  |> io.println
}
