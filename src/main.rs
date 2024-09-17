// use rayon::prelude::*;
use clap::{Parser, Subcommand};
use std::path::PathBuf;

mod animalhash;
mod boxtext;
mod cat_with_newline;
mod countext;
mod kg_for_bmi;
mod seqname;
mod serve;
mod setex;
mod week;

#[macro_use]
extern crate rocket;

#[derive(Debug, Parser)]
#[command(name = "swiss", about = "General utilities")]
struct Cli {
    #[command(subcommand)]
    command: Command,
}

#[derive(Subcommand, Debug, Clone)]
#[command(rename_all = "camel")]
enum Command {
    /// Basic web server
    #[command(alias = "s")]
    Serve,
    /// Wrap text in ASCII box art
    Boxtext { text: Vec<String> },
    /// Create new file by joining given files with space in between
    #[command(alias = "catNewline")]
    Join {
        filename: String,
        files: Vec<PathBuf>,
    },
    /// Output how many Kg required for a given BMI
    #[command(alias = "bmi")]
    KgForBMI { bmi: f32 },
    /// Rename to a sequential number, with optional prefix and suffix
    Seqname {
        /// Text to insert before number
        #[arg(short, long, default_value = "")]
        prefix: String,

        /// Text to insert after number
        #[arg(short, long, default_value = "")]
        suffix: String,

        ///Keep current name while prefixing or suffixing
        #[arg(short, long)]
        keep_filename: bool,

        /// Show files moved/renamed
        #[arg(short, long)]
        verbose: bool,

        /// Show files moved/renamed
        #[arg(short, long)]
        dry_run: bool,

        /// Separator for file components
        #[arg(long, default_value = "")]
        separator: String,

        /// Directories with files to rename
        dirs: Vec<PathBuf>,
    },
    /// Turn text into setex or atx header, depending on level
    Setex { level: i8, text: Vec<String> },
    /// Display start (Monday) and end (Sunday) of current week
    Week,
    /// Display start (Monday)
    WeekStart,
    /// Count Extensions
    CountExt,
    /// Generate AnimalHash-type phrase
    #[command(alias = "ah")]
    AnimalHash {
        /// --no-adjective     Don't include adjective
        #[arg(long, default_value = "true")]
        adjective: bool,
        ///    --no-animal        Don't include animal
        #[arg(long, default_value = "true")]
        animal: bool,
        ///    --no-colour        Don't include colour
        #[arg(long, default_value = "true")]
        colour: bool,
        ///    --camelcase        Use camelcase instead of kebabcase
        #[arg(long, default_value = "true")]
        camelcase: bool,
        ///    --pascalcase       Use pascalcase instead of kebabcase (capitalise every word)
        #[arg(long, default_value = "false")]
        pascalcase: bool,
    },
}

fn main() {
    let opts = Cli::parse();

    if let Err(e) = match opts.command.clone() {
        Command::Boxtext { text } => boxtext::run(text.join(" ")),
        Command::Join { filename, files } => cat_with_newline::run(&files, filename),
        Command::KgForBMI { bmi } => kg_for_bmi::run(bmi),
        Command::Seqname {
            prefix,
            suffix,
            keep_filename,
            verbose,
            dry_run,
            separator,
            dirs,
        } => seqname::run(
            prefix,
            suffix,
            keep_filename,
            verbose,
            dry_run,
            separator,
            &dirs,
        ),
        Command::Serve => serve::run(),
        Command::Setex { level, text } => setex::run(level, &text),
        Command::Week => week::start_and_end(),
        Command::WeekStart => week::start(),
        Command::CountExt => countext::run(),
        Command::AnimalHash {
            adjective,
            animal,
            colour,
            camelcase,
            pascalcase,
        } => animalhash::run(adjective, animal, colour, camelcase, pascalcase),
    } {
        eprintln!("{:?}, {}", opts.command, e);
    }
}
