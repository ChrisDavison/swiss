// use rayon::prelude::*;
use std::path::PathBuf;
use structopt::StructOpt;

mod boxtext;
mod cat_with_newline;
mod kg_for_bmi;
mod seqname;
mod serve;
mod setex;
mod week;

#[macro_use]
extern crate rocket;

#[derive(Debug, StructOpt)]
#[structopt(name = "swiss", about = "General utilities")]
struct Opts {
    #[structopt(subcommand)]
    command: Subcommand,
    // /// Use JSON rather than plaintext output
    // #[structopt(long, short)]
    // json: bool,
}

#[derive(Debug, StructOpt, PartialEq, Clone)]
enum Subcommand {
    /// Basic web server
    #[structopt(alias = "s")]
    Serve,
    /// Wrap text in ASCII box art
    Boxtext { text: Vec<String> },
    /// Create new file by joining given files with space in between
    CatWithNewline {
        filename: String,
        files: Vec<PathBuf>,
    },
    /// Output how many Kg required for a given BMI
    KgForBMI { bmi: f32 },
    /// Rename to a sequential number, with optional prefix and suffix
    Seqname {
        /// Text to insert before number
        #[structopt(short, long, default_value = "")]
        prefix: String,

        /// Text to insert after number
        #[structopt(short, long, default_value = "")]
        suffix: String,

        ///Keep current name while prefixing or suffixing
        #[structopt(short, long)]
        keep_filename: bool,

        /// Show files moved/renamed
        #[structopt(short, long)]
        verbose: bool,

        /// Show files moved/renamed
        #[structopt(short, long)]
        dry_run: bool,

        /// Separator for file components
        #[structopt(long, default_value = "")]
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
}

fn main() {
    let opts = Opts::from_args();

    if let Err(e) = match opts.command.clone() {
        Subcommand::Boxtext { text } => boxtext::run(text.join(" ")),
        Subcommand::CatWithNewline { filename, files } => cat_with_newline::run(&files, filename),
        Subcommand::KgForBMI { bmi } => kg_for_bmi::run(bmi),
        Subcommand::Seqname {
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
        Subcommand::Serve => serve::run(),
        Subcommand::Setex { level, text } => setex::run(level, &text),
        Subcommand::Week => week::start_and_end(),
        Subcommand::WeekStart => week::start(),
    } {
        eprintln!("{:?}, {}", opts.command, e);
    }
}
