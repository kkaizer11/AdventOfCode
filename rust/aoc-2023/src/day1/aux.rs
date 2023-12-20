use std::fs::File;
use std::io::prelude::*;

pub fn diga_oi() {
    println!("Dizendo oi");
}

pub fn read() -> std::io::Result<()> {
    let mut file = File::open("day1/input.txt")?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    assert_eq!(contents, "Hello, world!");
    Ok(())
}
