import os
import random
import string

# Directory where files will be created
output_dir = "files_output"
os.makedirs(output_dir, exist_ok=True)

# Function to generate filenames matching the regex 'phan[a-z][a-z][0-9][0-9]'
def generate_matching_name():
    letters = ''.join(random.choices(string.ascii_lowercase, k=2))
    numbers = ''.join(random.choices(string.digits, k=2))
    return f"phan{letters}{numbers}"

# Function to generate random filenames not matching the regex
def generate_random_name():
    length = random.randint(5, 12)
    return ''.join(random.choices(string.ascii_letters + string.digits, k=length))

# Create 12 files following the regex pattern
for _ in range(12):
    filename = generate_matching_name() + ".txt"
    with open(os.path.join(output_dir, filename), 'w') as file:
        file.write("This file matches the regex pattern.")

# Create 38 files that do not match the regex
for _ in range(38):
    filename = generate_random_name() + ".txt"
    with open(os.path.join(output_dir, filename), 'w') as file:
        file.write("This file does not match the regex pattern.")

print(f"50 files have been created in the '{output_dir}' directory.")
