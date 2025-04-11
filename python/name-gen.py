#!/usr/bin/env python3
import argparse
import random
import json
import os

# Expand the tilde to the full path
name_options_path = os.path.expanduser("~/my-code/name-options.json")

# Load the JSON file once
with open(name_options_path, "r") as f:
    name_options = json.load(f)


# Function to roll two six-sided dice
def roll_dice_key() -> str:
    return f"{random.randint(1, 6)}{random.randint(1, 6)}"


def safe_get(table: dict, key: str, category: str) -> str:
    try:
        return table[key]
    except KeyError:
        raise KeyError(f"Missing key '{key}' in category '{category}'.")


# General function to generate a full name
def generate_name(name_type: str) -> str:
    prefixes = name_options[name_type]["prefixes"]
    suffixes = name_options[name_type]["suffixes"]
    last_prefixes = name_options["last_names"]["prefixes"]
    last_suffixes = name_options["last_names"]["suffixes"]

    # Randomly select keys for first and last name parts
    first_roll_1 = roll_dice_key()
    first_roll_2 = roll_dice_key()
    last_roll_1 = roll_dice_key()
    last_roll_2 = roll_dice_key()

    first_name = safe_get(prefixes, first_roll_1, f"{name_type}.prefixes") + safe_get(
        suffixes, first_roll_2, f"{name_type}.suffixes"
    )
    last_name = safe_get(last_prefixes, last_roll_1, "last_names.prefixes") + safe_get(
        last_suffixes, last_roll_2, "last_names.suffixes"
    )

    return f"{first_name} {last_name}"


# Function to generate a male name
def generate_male_name() -> str:
    return generate_name("male_names")


# Function to generate a female name
def generate_female_name() -> str:
    return generate_name("female_names")


# Generate and print names
if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("gender", choices=["male", "female", "both"])
    args = parser.parse_args()
    if args.gender == "male":
        print(f"M: {generate_male_name()}")
    elif args.gender == "female":
        print(f"F: {generate_female_name()}")
    else:
        print(f"M: {generate_male_name()} F: {generate_female_name()}")
