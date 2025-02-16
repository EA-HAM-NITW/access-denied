#!/bin/bash

# Read credentials securely (hide password input)
echo "Enter admin username:"
read admin
echo "Enter team name:"
read team_name
echo "Enter team password:" 
read -s team_pwd  # -s flag hides password input

# Create new user with home directory and bash shell
sudo useradd -m -s /bin/bash "$team_name"

# Set password securely
echo "$team_name:$team_pwd" | sudo chpasswd


# Set ownership for all team's directories and files
sudo chown -R $admin:$team_name /home/$team_name

# Set permissions:
# For directories - admin(rwx), team(r-x), others(---)
sudo find /home/$team_name -type d -exec chmod 750 {} +

# For files - admin(rw-), team(r--), others(---)
sudo find /home/$team_name -type f -exec chmod 640 {} +

# Set admin's home directory permissions - no access for anyone else
sudo chmod -R 700 /home/$admin

# Create 404 directory
sudo mkdir -p /home/$team_name/404
sudo mkdir -p /home/$team_name/404/001

# Copy dataset with preserved permissions
sudo cp datasets/pincode.csv /home/$team_name/404/001/
sudo chown $admin:$team_name /home/$team_name/404/001/pincode.csv
sudo chmod 640 /home/$team_name/404/001/pincode.csv

sudo touch /home/$team_name/404/001/script.sh
sudo chown $admin:$team_name /home/$team_name/404/001/script.sh
sudo chmod 777 /home/$team_name/404/001/script.sh
