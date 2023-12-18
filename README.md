# Lupus 🌸
<p align="center">
      <img src="images/logo_lupus.png" alt="Image" width="400" />
</p>

[![Go version](https://img.shields.io/badge/Go-1.21.3%2B-blue)](https://golang.org/dl/)

# Description
Lupus is a small platformer for Yandex360, presented in binary format.

Main functions:

      * Logging of all events related to the work of documents
      * Notification of critical events
      * Create files to store information for up to 6 months
      * Data format - json and CEF syslog
      * Notification of new types of events that may subsequently appear
      * Lupus collects events for the last hour, so you need to set the time on the OS in the same way as on Yandex360


![lupus.png](images/lupus.png)

# How to work
---You need to select a binary file for the operating system on which you will run it (the binaries are in the cmd folder)---

1. For the correct logic of work, it is necessary to create any directory where you will launch the framework and move the binary there
2. Create directories logs, syslog and a hidden file with credentials for successful operation of .env
3. Give system rights to call this binary file

Example of organization on Linux

```console
sudo timedatectl set-timezone <(Continent)/(City)>
mkdir -p programs/{logs,syslogs} && touch programs/.env
cd programs
chmod +x lupus
```

!!! To work with the framework, you just need to download the binary, but you will need to fill the .env file with your credentials, the example is in the .env_example file

!!! Lupus has the ability to send notifications to Telegram, to do this you need to add a flag (-r/--report)

    Example: ./lupus --report

## Crontab

Sign up to run the framework every hour

```console
0 * * * * cd ~/programs && ./lupus --report
```



