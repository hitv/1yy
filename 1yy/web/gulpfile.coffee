gulp = require 'gulp'
gutil = require 'gulp-util'
plumber = require 'gulp-plumber'
runSequence = require('run-sequence').use gulp

less = require 'gulp-less'
coffee = require 'gulp-coffee'
concat = require 'gulp-concat'

minifier = require 'gulp-minifier'

FRONTEND_PATH = 'frontend'
PUBLIC_PATH = 'public'

errHandler = (err) ->
  gutil.beep()
  gutil.log err
  this.emit 'end'

# options
lessOption =
  paths: []

coffeeOption =
  bare: true

minifierOption =
  minify: true
  collapseWhitespace: true
  conservativeCollapse: true
  minifyJS: true
  minifyCSS: true

IS_PROD = false

COFFEE_SRC = [
  'coffee/modules/*.coffee'
  'coffee/*.coffee'
]

LIBS_SRC = [
  'lib/zepto.js'
  'lib/iscroll.js'
]

LESS_SRC = [
  'less/*.less'
]


JS_DEST_SRC_FILE = "js/main.src.js"
LIBS_DEST_SRC_FILE = 'js/libs.src.js'
CSS_DEST_SRC_FILE = 'css/main.src.css'

FINAL_JS_FILES = [
  LIBS_DEST_SRC_FILE,
  JS_DEST_SRC_FILE
]

FINAL_JS_SRC_FILE = "js/all.src.js"
FINAL_JS_MIN_FILE = "js/all.min.js"

gulp.task 'build-lib-js', ->
  gulp.src LIBS_SRC, cwd: FRONTEND_PATH
    .pipe plumber errorHandler: errHandler
    .pipe concat LIBS_DEST_SRC_FILE
    .pipe gulp.dest PUBLIC_PATH

gulp.task 'build-main-js', ->
  gulp.src COFFEE_SRC, cwd: FRONTEND_PATH
    .pipe plumber errorHandler: errHandler
    .pipe coffee coffeeOption
    .pipe concat JS_DEST_SRC_FILE
    .pipe gulp.dest PUBLIC_PATH

gulp.task 'contact-js', ->
  gulp.src FINAL_JS_FILES, cwd: PUBLIC_PATH
    .pipe plumber errorHandler: errHandler
    .pipe concat FINAL_JS_SRC_FILE
    .pipe gulp.dest PUBLIC_PATH

gulp.task 'build-css', ->
  gulp.src LESS_SRC, cwd: FRONTEND_PATH
    .pipe plumber errorHandler: errHandler
    .pipe less lessOption
    .pipe concat CSS_DEST_SRC_FILE, newLine: '\n'
    .pipe gulp.dest PUBLIC_PATH

gulp.task 'build', ['build-lib-js', 'build-main-js', 'build-css']

gulp.task 'all', ->
  runSequence 'build', 'contact-js'

gulp.task 'default', ['all']
