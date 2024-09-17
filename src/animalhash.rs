use rand::random;

fn title_case(word: &str) -> String {
    let lower = word.to_lowercase();
    let first = lower.chars().take(1).collect::<String>().to_uppercase();
    first + &lower[1..]
}

fn rand_line_from_string(string: &str) -> String {
    let lines: Vec<&str> = string.split('\n').collect();
    lines[random::<usize>() % lines.len()].trim().into()
}

pub fn run(
    adjective: bool,
    animal: bool,
    colour: bool,
    camelcase: bool,
    pascalcase: bool,
) -> Result<(), Box<dyn std::error::Error>> {
    let animals = include_str!("../assets/animals.txt");
    let adjectives = include_str!("../assets/adjectives.txt");
    let colours = include_str!("../assets/colours.txt");

    let mut outparts: Vec<String> = Vec::new();
    if adjective {
        outparts.push(rand_line_from_string(adjectives));
    }

    if colour {
        outparts.push(rand_line_from_string(colours));
    }

    if animal {
        outparts.push(rand_line_from_string(animals));
    }
    let mut joiner = "-";

    if camelcase || pascalcase {
        for elem in outparts.iter_mut().skip(if pascalcase { 0 } else { 1 }) {
            *elem = title_case(elem);
        }
        joiner = "";
    }
    let outstr = outparts.join(joiner);

    println!("{}", outstr);
    Ok(())
}
