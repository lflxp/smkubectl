compdef _demo demo

_demo() {
  typeset -A opt_args

  _arguments -C \
    '1:cmd:->cmds' \
    '2:generators:->generator_lists' \
    '-m[music file]:filename:->files' \
    '*:: :->args' \
  && ret=0

  case "$state" in
    (files)
        local -a music_files
        music_files=( ~/.music/**/*.{mp3,wav,flac,ogg} )
        _multi_parts / music_files
      ;;
    (cmds)
      local commands; commands=(
        "console:Boots up the Padrino application irb console (alternatively use 'c')"
        "generate:Executes the Padrino generator with given options (alternatively use 'gen' or 'g')"
        "help:Describe available commands or one specific command (alternatively use 'h')"
        "rake:Execute rake tasks"
        "runner:Run a piece of code in the Padrino application environment (alternatively use 'run' or 'r')"
        "start:Starts the Padrino application (alternatively use 's')"
        "stop:Stops the Padrino application (alternatively use 'st')"
        "version: Show current Padrino version."
        "create:Create a resource from a file or from stdin"
        "expose:Take a replication controller, service, deployment or pod and expose it as a new Kubernetes service"
        "run:Run a particular image on the cluster"
        "set:Set specific features on objects"

        "explain:Get documentation for a resource"
        "get:Display one or many resources"
        "edit:Edit a resource on the server"
        "delete:Delete resources by file names, stdin, resources and names, or by resources and label selector"

        "rollout:Manage the rollout of a resource"
        "scale:Set a new size for a deployment, replica set, or replication controller"
        "autoscale:Auto-scale a deployment, replica set, stateful set, or replication controller"

        "certificate:Modify certificate resources."
        "cluster-info:Display cluster information"
        "top:Display resource (CPU/memory) usage"
        "cordon:Mark node as unschedulable"
        "uncordon:Mark node as schedulable"
        "drain:Drain node in preparation for maintenance"
        "taint:Update the taints on one or more nodes"

        "describe:Show details of a specific resource or group of resources"
        "logs:Print the logs for a container in a pod"
        "attach:Attach to a running container"
        "exec:Execute a command in a container"
        "port-forward:Forward one or more local ports to a pod"
        "proxy:Run a proxy to the Kubernetes API server"
        "cp:Copy files and directories to and from containers"
        "auth:Inspect authorization"
        "debug:Create debugging sessions for troubleshooting workloads and nodes"

        "diff:Diff the live version against a would-be applied version"
        "apply:Apply a configuration to a resource by file name or stdin"
        "patch:Update fields of a resource"
        "replace:Replace a resource by file name or stdin"
        "wait:Experimental: Wait for a specific condition on one or many resources"
        "kustomize:Build a kustomization target from a directory or URL."

        "label:Update the labels on a resource"
        "annotate:Update the annotations on a resource"
        "completion:Output shell completion code for the specified shell (bash or zsh)"

        "api-resources:Print the supported API resources on the server"
        "api-versions:Print the supported API versions on the server, in the form of 'group/version'"
        "config:Modify kubeconfig files"
        "plugin:Provides utilities for interacting with plugins"
        "version:Print the client and server version information"
      )

      _describe -t commands 'command' commands && ret=0

      case $line[1] in
        (get)
          local get; gets=(
            "pod: pod"
            "svc: svc"
            "deploy: deploy"
          )
        ;;
      esac
    ;;
    (generator_lists)
      case $line[1] in
        (generate)
          local generate; generates=(
            "app:generates a new Padrino application"
            "controller:generates a new Padrino controller"
            "component:add components into a Padrino project"
            "helper:generates a new Padrino helper"
            "mailer:generates a new Padrino mailer"
            "migration:generates a new migration file"
            "model:generates a new model and migration files"
            "project:generates a new Padrino project"
            "plugin:sets up a plugin within a Padrino application"
            "task:generates a new task file"
          )
          _describe -t generates 'generate' generates && ret=0
        ;;
        (start)
          local start; starts=(
          "-p:specify the port for the running Padrino instance"
          )
          _describe -t starts 'start' starts && ret=0
        ;;
        (rake)
          local rake; rakes=(
            'ar\:abort_if_pending_migrations:Raises an error if there are pending migrations.'
            'ar\:charset:Retrieves the charset for the current Padrino.env database.'
            'ar\:collation:Retrieves the collation for the current Padrino.env database.'
            'ar\:create:Creates the database as defined in config/database.yml'
            'ar\:create\:all:Creates local databases as defined in config/database.yml'
            'ar\:drop:Drops the database for the current Padrino.env'
            'ar\:drop\:all:Drops local databases defined in config/database.yml'
            'ar\:forward:Pushes the schema to the next version.'
            'ar\:migrate:Migrates the database through scripts in db/migrate'
            'ar\:migrate\:down:Runs the "down" for a given migration VERSION.'
            'ar\:migrate\:redo:Rollbacks the database one migration and re migrate up.'
            'ar\:migrate\:reset:Resets your database using your migrations for current Padrino.env.'
            'ar\:migrate\:up:Runs the "up" for a given migration VERSION NUMBER.'
            'ar\:reset:Drops and recreates the database using db/schema.rb'
            'ar\:rollback:Rolls back the schema to previous schema version.'
            'ar\:schema\:dump:Creates a portable db/schema.rb file.'
            'ar\:schema\:load:Loads a schema.rb file into the database.'
            'ar\:setup:Creates the database, loads the schema, and seeds data.'
            'ar\:structure\:dump:Dumps the database structure to a SQL file.'
            'ar\:translate:Translate your ActiveRecord models.'
            'ar\:version:Retrieves the current schema version number.'
            'dm\:auto\:migrate:Performs an automigration (resets your db data).'
            'dm\:auto\:upgrade:Performs a non destructive automigration.'
            'dm\:create:Creates the database.'
            'dm\:drop:Drops the database (postgres and mysql only).'
            'dm\:migrate:Migrates the database to the latest version.'
            'dm\:migrate\:down[version]:Migrates down using migrations.'
            'dm\:migrate\:up[version]:Migrates up using migrations.'
            'dm\:reset:Drops the database, and migrates from scratch.'
            'dm\:setup:Create the database migrate and initialize with the seed data.'
            'sq\:migrate\:auto:Perform automigration (reset your db data).'
            'sq\:migrate\:to[version]:Perform migration up/down to VERSION.'
            'sq\:migrate\:up:Perform migration up to latest migration available.'
            'sq\:migrate\:down:Perform migration down (erase all data).'
            'mm\:translate:Generates .yml files for I18n translations.'
            'mi\:drop:Drops all the collections for the database for the current environment.'
            'mi\:create_indexes:Create the indexes defined on your mongoid models.'
            'mi\:objectid_convert:Convert string objectids in mongo database to ObjectID type.'
            'mi\:cleanup_old_collections:Clean up old collections backed up by objectid_convert.'
            'locale\:models:Generate a YAML file for localizing your models.'
            'routes:Retrieving a list of named routes.'
            'spec:Run all rspec tests.'
            'test:Run all tests written in bacon, riot, and shoulda.'
          )
          _describe -t rakes 'rake' rakes && ret=0
        ;;
      esac
    ;;
    (args)
      case $line[2] in
        (app)
          local apps; apps=(
            '-f:Generate app files if app already exists'
            '--force:Generate app files if app already exists'
          )
          _describe -t apps 'app' apps && ret=0
        ;;
        (controller)
          local controllers; controllers=(
            '-d:remove all generated files'
            '--destroy:remove all generated files'
            '-n:The name space of your Padrino project'
            '--namespace:The name space of your Padrino project'
            '-l:The layout for the controller'
            '--layout:The layout for the controller'
            '-p:The parent of the controller'
            '--parent:The parent of the controller'
            '-r:The root destination'
            '--root:The root destination'
            '--no-helper:Not generate helper'
          )
          _describe -t controllers 'controller' controllers && ret=0
        ;;
        (component)
          local components; components=(
            '-n: The application name'
            '--app==APP: The application name'
            '-r: The root destination => Default: .'
            '--root=ROOT: The root destination => Default: .'
            '-a: SQL adapter for ORM (sqlite, mysql, mysql2, mysql-gem, postgres) => Default: sqlite'
            '--adapter=ADAPTER: SQL adapter for ORM (sqlite, mysql, mysql2, mysql-gem, postgres) => Default: sqlite'
            '-d: The database engine component (mongomatic, sequel, minirecord, dynamoid, activerecord, ohm, mongomapper, couchrest, ripple, mongoid, datamapper, none)'
            '--orm=ORM: The database engine component (mongomatic, sequel, minirecord, dynamoid, activerecord, ohm, mongomapper, couchrest, ripple, mongoid, datamapper, none)'
            '-t: The testing framework component (minitest, cucumber, shoulda, steak, riot, bacon, rspec, none)'
            '--test=TEST: The testing framework component (minitest, cucumber, shoulda, steak, riot, bacon, rspec, none)'
            '-m: The mocking library component (mocha, rr, none)'
            '--mock=MOCK: The mocking library component (mocha, rr, none)'
            '-s: The javascript library component (jquery, extcore, dojo, rightjs, mootools, prototype, none)'
            '--script=SCRIPT: The javascript library component (jquery, extcore, dojo, rightjs, mootools, prototype, none)'
            '-e: The template engine component (liquid, haml, slim, erb, none)'
            '--renderer=RENDERER: The template engine component (liquid, haml, slim, erb, none)'
            '-c: The stylesheet engine component (compass, less, scss, sass, none)'
            '--stylesheet=STYLESHEET: The stylesheet engine component (compass, less, scss, sass, none)'
          )
          _describe -t components 'component' components && ret=0
        ;;
        (mailer)
          local mailers; mailers=(
            '-a:creates a mailer for the specified subapp'
          )
          _describe -t mailers 'mailer' mailers && ret=0
        ;;
        (migration)
          local migrations; migrations=(
            '-d:removes all generated files'
            '-r:specify the root destination path'
          )
          _describe -t migrations 'migration' migrations && ret=0
        ;;
        (model)
          local models; models=(
            '-d:removes all generated files'
            '-r:specify the root destination path'
            '-s:skip migration generation'
            '-f:Generate app files if app already exists'
            '--force:Generate app files if app already exists'
          )
          _describe -t models 'model' models && ret=0
        ;;
        (project)
          local projects; projects=(
            "-a:specify db adapter (options: 'mysql', 'mysql2', 'mysql-gem', sqlite' , 'postgres')"
            "-b:execute bundler dependencies installation"
            "-c:define stylesheet (options: 'compass', 'less', 'sass', 'scss', 'none')"
            "-d:define orm (options: 'activerecord', 'couchrest', 'mongoid', 'datamapper', 'minirecord', 'mongomatic', 'mongomapper', 'ohm', 'ripple', 'sequel', 'none')"
            "-e:define renderer (options: 'erb', 'haml','liquid', 'slim', 'none')"
            "-i:generate tiny project skeleton without any components"
            "-m:define mock (options: 'mocha', 'rr', 'none')"
            "-n:specify app name different from the project name"
            "-r:the root destination path for the project"
            "-s:define script (options: 'dojo', 'extcore', 'jquery', 'mootools', 'prototype', 'rightjs', 'none')"
            "-t:define test (options: 'bacon', 'cucumber', 'shoulda', 'riot', 'rspec', 'steak', 'minitest', 'none')"
          )
          _describe -t projects 'project' projects && ret=0
        ;;
      esac
    ;;
  esac

  return 1
}
