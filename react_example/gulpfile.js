var gulp = require('gulp');
var browserify = require('browserify');
var babelify = require('babelify');
var source = require('vinyl-source-stream');

function build() {
   return browserify({entries: 'src/app.jsx', extensions: ['.jsx'], debug: false})
        .transform('babelify', {presets: ['es2015', 'react']})
        .transform(require('browserify-css'))
        .bundle()
        .pipe(source('app.js'))
        .pipe(gulp.dest('./dist'));
}

function watch(module) {
    gulp.watch('src/*', ['build']);
}

gulp.task('build', build);
gulp.task('watch', watch);
gulp.task('default', ['build', 'watch']);
