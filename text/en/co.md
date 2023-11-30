Print out git co-author line

Co-author is a unix filter intended to help users write co-author lines when committing from the command line.

To start, use the command {{aka}} {key}

Prerequisites The command requires a yaml config to be specified as the following:

co_authors:
  jeffs: # Unique identifier for Jeff Smith
    full_name: Jeff Smith
    email: jeff.smith@example.com
  jeffst: # Unique identifier for Jeff Street
    full_name: Jeff Street
    email: jeff.street@example.com
  sarahj:
    full_name: Sarah Jones
    email: sarah.jones@example.com
  alexb:
    full_name: Alex Brown
    email: alex.brown@example.com
  # Add more co-authors with unique identifiers
