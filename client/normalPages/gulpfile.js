var gulp = require('gulp');
var babel = require("gulp-babel");
var autoprefixer = require('gulp-autoprefixer');
var px2rem = require('gulp-px2rem');
gulp.task('js', function () {
  return gulp.src("js/*.js")
    .pipe(babel())
    .pipe(gulp.dest("dist/js"))
})

gulp.task('css', function () {
  return gulp.src("css/*.css")
    .pipe(px2rem({
      rootValue: 150,
      unitPrecision: 5
    }))
    .pipe(autoprefixer())
    .pipe(gulp.dest("dist/css"))
})
gulp.task('html', function () {
  return gulp.src("*.html")
    .pipe(gulp.dest("dist"))
})

gulp.task('default', ['js','css', 'html'], function () {
  // body...
})