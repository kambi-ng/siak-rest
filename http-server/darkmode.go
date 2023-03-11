package httpserver

/*
@name        Amoenus Swagger Dark
@namespace   https://github.com/Amoenus/SwaggerDark
@license     MIT
@author      Amoenus
*/

var darkModeCss = `
a { color: #8c8cfa; }

::-webkit-scrollbar-track-piece { background-color: rgba(255, 255, 255, .2) !important; }

::-webkit-scrollbar-track { background-color: rgba(255, 255, 255, .3) !important; }

::-webkit-scrollbar-thumb { background-color: rgba(255, 255, 255, .5) !important; }

embed[type="application/pdf"] { filter: invert(90%); }

html {
    background: #1f1f1f !important;
    box-sizing: border-box;
    filter: contrast(100%) brightness(100%) saturate(100%);
    overflow-y: scroll;
}

body {
    background: #1f1f1f;
    background-color: #1f1f1f;
    background-image: none !important;
}

button, input, select, textarea {
    background-color: #1f1f1f;
    color: #bfbfbf;
}

font, html { color: #bfbfbf; }

.swagger-ui, .swagger-ui section h3 { color: #b5bac9; }

.swagger-ui a { background-color: transparent; }

.swagger-ui mark {
    background-color: #664b00;
    color: #bfbfbf;
}

.swagger-ui legend { color: inherit; }

.swagger-ui .debug * { outline: #e6da99 solid 1px; }

.swagger-ui .debug-white * { outline: #fff solid 1px; }

.swagger-ui .debug-black * { outline: #bfbfbf solid 1px; }

.swagger-ui .debug-grid { background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAYAAADED76LAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyhpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTExIDc5LjE1ODMyNSwgMjAxNS8wOS8xMC0wMToxMDoyMCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6MTRDOTY4N0U2N0VFMTFFNjg2MzZDQjkwNkQ4MjgwMEIiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6MTRDOTY4N0Q2N0VFMTFFNjg2MzZDQjkwNkQ4MjgwMEIiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUgKE1hY2ludG9zaCkiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDo3NjcyQkQ3NjY3QzUxMUU2QjJCQ0UyNDA4MTAwMjE3MSIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDo3NjcyQkQ3NzY3QzUxMUU2QjJCQ0UyNDA4MTAwMjE3MSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/PsBS+GMAAAAjSURBVHjaYvz//z8DLsD4gcGXiYEAGBIKGBne//fFpwAgwAB98AaF2pjlUQAAAABJRU5ErkJggg==) 0 0; }

.swagger-ui .debug-grid-16 { background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyhpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTExIDc5LjE1ODMyNSwgMjAxNS8wOS8xMC0wMToxMDoyMCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6ODYyRjhERDU2N0YyMTFFNjg2MzZDQjkwNkQ4MjgwMEIiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6ODYyRjhERDQ2N0YyMTFFNjg2MzZDQjkwNkQ4MjgwMEIiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUgKE1hY2ludG9zaCkiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDo3NjcyQkQ3QTY3QzUxMUU2QjJCQ0UyNDA4MTAwMjE3MSIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDo3NjcyQkQ3QjY3QzUxMUU2QjJCQ0UyNDA4MTAwMjE3MSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/PvCS01IAAABMSURBVHjaYmR4/5+BFPBfAMFm/MBgx8RAGWCn1AAmSg34Q6kBDKMGMDCwICeMIemF/5QawEipAWwUhwEjMDvbAWlWkvVBwu8vQIABAEwBCph8U6c0AAAAAElFTkSuQmCC) 0 0; }

.swagger-ui .debug-grid-8-solid { background: url(data:image/jpeg;base64,/9j/4QAYRXhpZgAASUkqAAgAAAAAAAAAAAAAAP/sABFEdWNreQABAAQAAAAAAAD/4QMxaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLwA8P3hwYWNrZXQgYmVnaW49Iu+7vyIgaWQ9Ilc1TTBNcENlaGlIenJlU3pOVGN6a2M5ZCI/PiA8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIiB4OnhtcHRrPSJBZG9iZSBYTVAgQ29yZSA1LjYtYzExMSA3OS4xNTgzMjUsIDIwMTUvMDkvMTAtMDE6MTA6MjAgICAgICAgICI+IDxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+IDxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSIiIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bXA6Q3JlYXRvclRvb2w9IkFkb2JlIFBob3Rvc2hvcCBDQyAyMDE1IChNYWNpbnRvc2gpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOkIxMjI0OTczNjdCMzExRTZCMkJDRTI0MDgxMDAyMTcxIiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOkIxMjI0OTc0NjdCMzExRTZCMkJDRTI0MDgxMDAyMTcxIj4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6QjEyMjQ5NzE2N0IzMTFFNkIyQkNFMjQwODEwMDIxNzEiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6QjEyMjQ5NzI2N0IzMTFFNkIyQkNFMjQwODEwMDIxNzEiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz7/7gAOQWRvYmUAZMAAAAAB/9sAhAAbGhopHSlBJiZBQi8vL0JHPz4+P0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHAR0pKTQmND8oKD9HPzU/R0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0f/wAARCAAIAAgDASIAAhEBAxEB/8QAWQABAQAAAAAAAAAAAAAAAAAAAAYBAQEAAAAAAAAAAAAAAAAAAAIEEAEBAAMBAAAAAAAAAAAAAAABADECA0ERAAEDBQAAAAAAAAAAAAAAAAARITFBUWESIv/aAAwDAQACEQMRAD8AoOnTV1QTD7JJshP3vSM3P//Z) 0 0 #1c1c21; }

.swagger-ui .debug-grid-16-solid { background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAIAAACQkWg2AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyhpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTExIDc5LjE1ODMyNSwgMjAxNS8wOS8xMC0wMToxMDoyMCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUgKE1hY2ludG9zaCkiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6NzY3MkJEN0U2N0M1MTFFNkIyQkNFMjQwODEwMDIxNzEiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6NzY3MkJEN0Y2N0M1MTFFNkIyQkNFMjQwODEwMDIxNzEiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDo3NjcyQkQ3QzY3QzUxMUU2QjJCQ0UyNDA4MTAwMjE3MSIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDo3NjcyQkQ3RDY3QzUxMUU2QjJCQ0UyNDA4MTAwMjE3MSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/Pve6J3kAAAAzSURBVHjaYvz//z8D0UDsMwMjSRoYP5Gq4SPNbRjVMEQ1fCRDg+in/6+J1AJUxsgAEGAA31BAJMS0GYEAAAAASUVORK5CYII=) 0 0 #1c1c21; }

.swagger-ui .b--black { border-color: #000; }

.swagger-ui .b--near-black { border-color: #121212; }

.swagger-ui .b--dark-gray { border-color: #333; }

.swagger-ui .b--mid-gray { border-color: #545454; }

.swagger-ui .b--gray { border-color: #787878; }

.swagger-ui .b--silver { border-color: #999; }

.swagger-ui .b--light-silver { border-color: #6e6e6e; }

.swagger-ui .b--moon-gray { border-color: #4d4d4d; }

.swagger-ui .b--light-gray { border-color: #2b2b2b; }

.swagger-ui .b--near-white { border-color: #242424; }

.swagger-ui .b--white { border-color: #1c1c21; }

.swagger-ui .b--white-90 { border-color: rgba(28, 28, 33, .9); }

.swagger-ui .b--white-80 { border-color: rgba(28, 28, 33, .8); }

.swagger-ui .b--white-70 { border-color: rgba(28, 28, 33, .7); }

.swagger-ui .b--white-60 { border-color: rgba(28, 28, 33, .6); }

.swagger-ui .b--white-50 { border-color: rgba(28, 28, 33, .5); }

.swagger-ui .b--white-40 { border-color: rgba(28, 28, 33, .4); }

.swagger-ui .b--white-30 { border-color: rgba(28, 28, 33, .3); }

.swagger-ui .b--white-20 { border-color: rgba(28, 28, 33, .2); }

.swagger-ui .b--white-10 { border-color: rgba(28, 28, 33, .1); }

.swagger-ui .b--white-05 { border-color: rgba(28, 28, 33, .05); }

.swagger-ui .b--white-025 { border-color: rgba(28, 28, 33, .024); }

.swagger-ui .b--white-0125 { border-color: rgba(28, 28, 33, .01); }

.swagger-ui .b--black-90 { border-color: rgba(0, 0, 0, .9); }

.swagger-ui .b--black-80 { border-color: rgba(0, 0, 0, .8); }

.swagger-ui .b--black-70 { border-color: rgba(0, 0, 0, .7); }

.swagger-ui .b--black-60 { border-color: rgba(0, 0, 0, .6); }

.swagger-ui .b--black-50 { border-color: rgba(0, 0, 0, .5); }

.swagger-ui .b--black-40 { border-color: rgba(0, 0, 0, .4); }

.swagger-ui .b--black-30 { border-color: rgba(0, 0, 0, .3); }

.swagger-ui .b--black-20 { border-color: rgba(0, 0, 0, .2); }

.swagger-ui .b--black-10 { border-color: rgba(0, 0, 0, .1); }

.swagger-ui .b--black-05 { border-color: rgba(0, 0, 0, .05); }

.swagger-ui .b--black-025 { border-color: rgba(0, 0, 0, .024); }

.swagger-ui .b--black-0125 { border-color: rgba(0, 0, 0, .01); }

.swagger-ui .b--dark-red { border-color: #bc2f36; }

.swagger-ui .b--red { border-color: #c83932; }

.swagger-ui .b--light-red { border-color: #ab3c2b; }

.swagger-ui .b--orange { border-color: #cc6e33; }

.swagger-ui .b--purple { border-color: #5e2ca5; }

.swagger-ui .b--light-purple { border-color: #672caf; }

.swagger-ui .b--dark-pink { border-color: #ab2b81; }

.swagger-ui .b--hot-pink { border-color: #c03086; }

.swagger-ui .b--pink { border-color: #8f2464; }

.swagger-ui .b--light-pink { border-color: #721d4d; }

.swagger-ui .b--dark-green { border-color: #1c6e50; }

.swagger-ui .b--green { border-color: #279b70; }

.swagger-ui .b--light-green { border-color: #228762; }

.swagger-ui .b--navy { border-color: #0d1d35; }

.swagger-ui .b--dark-blue { border-color: #20497e; }

.swagger-ui .b--blue { border-color: #4380d0; }

.swagger-ui .b--light-blue { border-color: #20517e; }

.swagger-ui .b--lightest-blue { border-color: #143a52; }

.swagger-ui .b--washed-blue { border-color: #0c312d; }

.swagger-ui .b--washed-green { border-color: #0f3d2c; }

.swagger-ui .b--washed-red { border-color: #411010; }

.swagger-ui .b--transparent { border-color: transparent; }

.swagger-ui .b--gold, .swagger-ui .b--light-yellow, .swagger-ui .b--washed-yellow, .swagger-ui .b--yellow { border-color: #664b00; }

.swagger-ui .shadow-1 { box-shadow: rgba(0, 0, 0, .2) 0 0 4px 2px; }

.swagger-ui .shadow-2 { box-shadow: rgba(0, 0, 0, .2) 0 0 8px 2px; }

.swagger-ui .shadow-3 { box-shadow: rgba(0, 0, 0, .2) 2px 2px 4px 2px; }

.swagger-ui .shadow-4 { box-shadow: rgba(0, 0, 0, .2) 2px 2px 8px 0; }

.swagger-ui .shadow-5 { box-shadow: rgba(0, 0, 0, .2) 4px 4px 8px 0; }

@media screen and (min-width: 30em) {
    .swagger-ui .shadow-1-ns { box-shadow: rgba(0, 0, 0, .2) 0 0 4px 2px; }

    .swagger-ui .shadow-2-ns { box-shadow: rgba(0, 0, 0, .2) 0 0 8px 2px; }

    .swagger-ui .shadow-3-ns { box-shadow: rgba(0, 0, 0, .2) 2px 2px 4px 2px; }

    .swagger-ui .shadow-4-ns { box-shadow: rgba(0, 0, 0, .2) 2px 2px 8px 0; }

    .swagger-ui .shadow-5-ns { box-shadow: rgba(0, 0, 0, .2) 4px 4px 8px 0; }
}

@media screen and (max-width: 60em) and (min-width: 30em) {
    .swagger-ui .shadow-1-m { box-shadow: rgba(0, 0, 0, .2) 0 0 4px 2px; }

    .swagger-ui .shadow-2-m { box-shadow: rgba(0, 0, 0, .2) 0 0 8px 2px; }

    .swagger-ui .shadow-3-m { box-shadow: rgba(0, 0, 0, .2) 2px 2px 4px 2px; }

    .swagger-ui .shadow-4-m { box-shadow: rgba(0, 0, 0, .2) 2px 2px 8px 0; }

    .swagger-ui .shadow-5-m { box-shadow: rgba(0, 0, 0, .2) 4px 4px 8px 0; }
}

@media screen and (min-width: 60em) {
    .swagger-ui .shadow-1-l { box-shadow: rgba(0, 0, 0, .2) 0 0 4px 2px; }

    .swagger-ui .shadow-2-l { box-shadow: rgba(0, 0, 0, .2) 0 0 8px 2px; }

    .swagger-ui .shadow-3-l { box-shadow: rgba(0, 0, 0, .2) 2px 2px 4px 2px; }

    .swagger-ui .shadow-4-l { box-shadow: rgba(0, 0, 0, .2) 2px 2px 8px 0; }

    .swagger-ui .shadow-5-l { box-shadow: rgba(0, 0, 0, .2) 4px 4px 8px 0; }
}

.swagger-ui .black-05 { color: rgba(191, 191, 191, .05); }

.swagger-ui .bg-black-05 { background-color: rgba(0, 0, 0, .05); }

.swagger-ui .black-90, .swagger-ui .hover-black-90:focus, .swagger-ui .hover-black-90:hover { color: rgba(191, 191, 191, .9); }

.swagger-ui .black-80, .swagger-ui .hover-black-80:focus, .swagger-ui .hover-black-80:hover { color: rgba(191, 191, 191, .8); }

.swagger-ui .black-70, .swagger-ui .hover-black-70:focus, .swagger-ui .hover-black-70:hover { color: rgba(191, 191, 191, .7); }

.swagger-ui .black-60, .swagger-ui .hover-black-60:focus, .swagger-ui .hover-black-60:hover { color: rgba(191, 191, 191, .6); }

.swagger-ui .black-50, .swagger-ui .hover-black-50:focus, .swagger-ui .hover-black-50:hover { color: rgba(191, 191, 191, .5); }

.swagger-ui .black-40, .swagger-ui .hover-black-40:focus, .swagger-ui .hover-black-40:hover { color: rgba(191, 191, 191, .4); }

.swagger-ui .black-30, .swagger-ui .hover-black-30:focus, .swagger-ui .hover-black-30:hover { color: rgba(191, 191, 191, .3); }

.swagger-ui .black-20, .swagger-ui .hover-black-20:focus, .swagger-ui .hover-black-20:hover { color: rgba(191, 191, 191, .2); }

.swagger-ui .black-10, .swagger-ui .hover-black-10:focus, .swagger-ui .hover-black-10:hover { color: rgba(191, 191, 191, .1); }

.swagger-ui .hover-white-90:focus, .swagger-ui .hover-white-90:hover, .swagger-ui .white-90 { color: rgba(255, 255, 255, .9); }

.swagger-ui .hover-white-80:focus, .swagger-ui .hover-white-80:hover, .swagger-ui .white-80 { color: rgba(255, 255, 255, .8); }

.swagger-ui .hover-white-70:focus, .swagger-ui .hover-white-70:hover, .swagger-ui .white-70 { color: rgba(255, 255, 255, .7); }

.swagger-ui .hover-white-60:focus, .swagger-ui .hover-white-60:hover, .swagger-ui .white-60 { color: rgba(255, 255, 255, .6); }

.swagger-ui .hover-white-50:focus, .swagger-ui .hover-white-50:hover, .swagger-ui .white-50 { color: rgba(255, 255, 255, .5); }

.swagger-ui .hover-white-40:focus, .swagger-ui .hover-white-40:hover, .swagger-ui .white-40 { color: rgba(255, 255, 255, .4); }

.swagger-ui .hover-white-30:focus, .swagger-ui .hover-white-30:hover, .swagger-ui .white-30 { color: rgba(255, 255, 255, .3); }

.swagger-ui .hover-white-20:focus, .swagger-ui .hover-white-20:hover, .swagger-ui .white-20 { color: rgba(255, 255, 255, .2); }

.swagger-ui .hover-white-10:focus, .swagger-ui .hover-white-10:hover, .swagger-ui .white-10 { color: rgba(255, 255, 255, .1); }

.swagger-ui .hover-moon-gray:focus, .swagger-ui .hover-moon-gray:hover, .swagger-ui .moon-gray { color: #ccc; }

.swagger-ui .hover-light-gray:focus, .swagger-ui .hover-light-gray:hover, .swagger-ui .light-gray { color: #ededed; }

.swagger-ui .hover-near-white:focus, .swagger-ui .hover-near-white:hover, .swagger-ui .near-white { color: #f5f5f5; }

.swagger-ui .dark-red, .swagger-ui .hover-dark-red:focus, .swagger-ui .hover-dark-red:hover { color: #e6999d; }

.swagger-ui .hover-red:focus, .swagger-ui .hover-red:hover, .swagger-ui .red { color: #e69d99; }

.swagger-ui .hover-light-red:focus, .swagger-ui .hover-light-red:hover, .swagger-ui .light-red { color: #e6a399; }

.swagger-ui .hover-orange:focus, .swagger-ui .hover-orange:hover, .swagger-ui .orange { color: #e6b699; }

.swagger-ui .gold, .swagger-ui .hover-gold:focus, .swagger-ui .hover-gold:hover { color: #e6d099; }

.swagger-ui .hover-yellow:focus, .swagger-ui .hover-yellow:hover, .swagger-ui .yellow { color: #e6da99; }

.swagger-ui .hover-light-yellow:focus, .swagger-ui .hover-light-yellow:hover, .swagger-ui .light-yellow { color: #ede6b6; }

.swagger-ui .hover-purple:focus, .swagger-ui .hover-purple:hover, .swagger-ui .purple { color: #b99ae4; }

.swagger-ui .hover-light-purple:focus, .swagger-ui .hover-light-purple:hover, .swagger-ui .light-purple { color: #bb99e6; }

.swagger-ui .dark-pink, .swagger-ui .hover-dark-pink:focus, .swagger-ui .hover-dark-pink:hover { color: #e699cc; }

.swagger-ui .hot-pink, .swagger-ui .hover-hot-pink:focus, .swagger-ui .hover-hot-pink:hover, .swagger-ui .hover-pink:focus, .swagger-ui .hover-pink:hover, .swagger-ui .pink { color: #e699c7; }

.swagger-ui .hover-light-pink:focus, .swagger-ui .hover-light-pink:hover, .swagger-ui .light-pink { color: #edb6d5; }

.swagger-ui .dark-green, .swagger-ui .green, .swagger-ui .hover-dark-green:focus, .swagger-ui .hover-dark-green:hover, .swagger-ui .hover-green:focus, .swagger-ui .hover-green:hover { color: #99e6c9; }

.swagger-ui .hover-light-green:focus, .swagger-ui .hover-light-green:hover, .swagger-ui .light-green { color: #a1e8ce; }

.swagger-ui .hover-navy:focus, .swagger-ui .hover-navy:hover, .swagger-ui .navy { color: #99b8e6; }

.swagger-ui .blue, .swagger-ui .dark-blue, .swagger-ui .hover-blue:focus, .swagger-ui .hover-blue:hover, .swagger-ui .hover-dark-blue:focus, .swagger-ui .hover-dark-blue:hover { color: #99bae6; }

.swagger-ui .hover-light-blue:focus, .swagger-ui .hover-light-blue:hover, .swagger-ui .light-blue { color: #a9cbea; }

.swagger-ui .hover-lightest-blue:focus, .swagger-ui .hover-lightest-blue:hover, .swagger-ui .lightest-blue { color: #d6e9f5; }

.swagger-ui .hover-washed-blue:focus, .swagger-ui .hover-washed-blue:hover, .swagger-ui .washed-blue { color: #f7fdfc; }

.swagger-ui .hover-washed-green:focus, .swagger-ui .hover-washed-green:hover, .swagger-ui .washed-green { color: #ebfaf4; }

.swagger-ui .hover-washed-yellow:focus, .swagger-ui .hover-washed-yellow:hover, .swagger-ui .washed-yellow { color: #fbf9ef; }

.swagger-ui .hover-washed-red:focus, .swagger-ui .hover-washed-red:hover, .swagger-ui .washed-red { color: #f9e7e7; }

.swagger-ui .color-inherit, .swagger-ui .hover-inherit:focus, .swagger-ui .hover-inherit:hover { color: inherit; }

.swagger-ui .bg-black-90, .swagger-ui .hover-bg-black-90:focus, .swagger-ui .hover-bg-black-90:hover { background-color: rgba(0, 0, 0, .9); }

.swagger-ui .bg-black-80, .swagger-ui .hover-bg-black-80:focus, .swagger-ui .hover-bg-black-80:hover { background-color: rgba(0, 0, 0, .8); }

.swagger-ui .bg-black-70, .swagger-ui .hover-bg-black-70:focus, .swagger-ui .hover-bg-black-70:hover { background-color: rgba(0, 0, 0, .7); }

.swagger-ui .bg-black-60, .swagger-ui .hover-bg-black-60:focus, .swagger-ui .hover-bg-black-60:hover { background-color: rgba(0, 0, 0, .6); }

.swagger-ui .bg-black-50, .swagger-ui .hover-bg-black-50:focus, .swagger-ui .hover-bg-black-50:hover { background-color: rgba(0, 0, 0, .5); }

.swagger-ui .bg-black-40, .swagger-ui .hover-bg-black-40:focus, .swagger-ui .hover-bg-black-40:hover { background-color: rgba(0, 0, 0, .4); }

.swagger-ui .bg-black-30, .swagger-ui .hover-bg-black-30:focus, .swagger-ui .hover-bg-black-30:hover { background-color: rgba(0, 0, 0, .3); }

.swagger-ui .bg-black-20, .swagger-ui .hover-bg-black-20:focus, .swagger-ui .hover-bg-black-20:hover { background-color: rgba(0, 0, 0, .2); }

.swagger-ui .bg-white-90, .swagger-ui .hover-bg-white-90:focus, .swagger-ui .hover-bg-white-90:hover { background-color: rgba(28, 28, 33, .9); }

.swagger-ui .bg-white-80, .swagger-ui .hover-bg-white-80:focus, .swagger-ui .hover-bg-white-80:hover { background-color: rgba(28, 28, 33, .8); }

.swagger-ui .bg-white-70, .swagger-ui .hover-bg-white-70:focus, .swagger-ui .hover-bg-white-70:hover { background-color: rgba(28, 28, 33, .7); }

.swagger-ui .bg-white-60, .swagger-ui .hover-bg-white-60:focus, .swagger-ui .hover-bg-white-60:hover { background-color: rgba(28, 28, 33, .6); }

.swagger-ui .bg-white-50, .swagger-ui .hover-bg-white-50:focus, .swagger-ui .hover-bg-white-50:hover { background-color: rgba(28, 28, 33, .5); }

.swagger-ui .bg-white-40, .swagger-ui .hover-bg-white-40:focus, .swagger-ui .hover-bg-white-40:hover { background-color: rgba(28, 28, 33, .4); }

.swagger-ui .bg-white-30, .swagger-ui .hover-bg-white-30:focus, .swagger-ui .hover-bg-white-30:hover { background-color: rgba(28, 28, 33, .3); }

.swagger-ui .bg-white-20, .swagger-ui .hover-bg-white-20:focus, .swagger-ui .hover-bg-white-20:hover { background-color: rgba(28, 28, 33, .2); }

.swagger-ui .bg-black, .swagger-ui .hover-bg-black:focus, .swagger-ui .hover-bg-black:hover { background-color: #000; }

.swagger-ui .bg-near-black, .swagger-ui .hover-bg-near-black:focus, .swagger-ui .hover-bg-near-black:hover { background-color: #121212; }

.swagger-ui .bg-dark-gray, .swagger-ui .hover-bg-dark-gray:focus, .swagger-ui .hover-bg-dark-gray:hover { background-color: #333; }

.swagger-ui .bg-mid-gray, .swagger-ui .hover-bg-mid-gray:focus, .swagger-ui .hover-bg-mid-gray:hover { background-color: #545454; }

.swagger-ui .bg-gray, .swagger-ui .hover-bg-gray:focus, .swagger-ui .hover-bg-gray:hover { background-color: #787878; }

.swagger-ui .bg-silver, .swagger-ui .hover-bg-silver:focus, .swagger-ui .hover-bg-silver:hover { background-color: #999; }

.swagger-ui .bg-white, .swagger-ui .hover-bg-white:focus, .swagger-ui .hover-bg-white:hover { background-color: #1c1c21; }

.swagger-ui .bg-transparent, .swagger-ui .hover-bg-transparent:focus, .swagger-ui .hover-bg-transparent:hover { background-color: transparent; }

.swagger-ui .bg-dark-red, .swagger-ui .hover-bg-dark-red:focus, .swagger-ui .hover-bg-dark-red:hover { background-color: #bc2f36; }

.swagger-ui .bg-red, .swagger-ui .hover-bg-red:focus, .swagger-ui .hover-bg-red:hover { background-color: #c83932; }

.swagger-ui .bg-light-red, .swagger-ui .hover-bg-light-red:focus, .swagger-ui .hover-bg-light-red:hover { background-color: #ab3c2b; }

.swagger-ui .bg-orange, .swagger-ui .hover-bg-orange:focus, .swagger-ui .hover-bg-orange:hover { background-color: #cc6e33; }

.swagger-ui .bg-gold, .swagger-ui .bg-light-yellow, .swagger-ui .bg-washed-yellow, .swagger-ui .bg-yellow, .swagger-ui .hover-bg-gold:focus, .swagger-ui .hover-bg-gold:hover, .swagger-ui .hover-bg-light-yellow:focus, .swagger-ui .hover-bg-light-yellow:hover, .swagger-ui .hover-bg-washed-yellow:focus, .swagger-ui .hover-bg-washed-yellow:hover, .swagger-ui .hover-bg-yellow:focus, .swagger-ui .hover-bg-yellow:hover { background-color: #664b00; }

.swagger-ui .bg-purple, .swagger-ui .hover-bg-purple:focus, .swagger-ui .hover-bg-purple:hover { background-color: #5e2ca5; }

.swagger-ui .bg-light-purple, .swagger-ui .hover-bg-light-purple:focus, .swagger-ui .hover-bg-light-purple:hover { background-color: #672caf; }

.swagger-ui .bg-dark-pink, .swagger-ui .hover-bg-dark-pink:focus, .swagger-ui .hover-bg-dark-pink:hover { background-color: #ab2b81; }

.swagger-ui .bg-hot-pink, .swagger-ui .hover-bg-hot-pink:focus, .swagger-ui .hover-bg-hot-pink:hover { background-color: #c03086; }

.swagger-ui .bg-pink, .swagger-ui .hover-bg-pink:focus, .swagger-ui .hover-bg-pink:hover { background-color: #8f2464; }

.swagger-ui .bg-light-pink, .swagger-ui .hover-bg-light-pink:focus, .swagger-ui .hover-bg-light-pink:hover { background-color: #721d4d; }

.swagger-ui .bg-dark-green, .swagger-ui .hover-bg-dark-green:focus, .swagger-ui .hover-bg-dark-green:hover { background-color: #1c6e50; }

.swagger-ui .bg-green, .swagger-ui .hover-bg-green:focus, .swagger-ui .hover-bg-green:hover { background-color: #279b70; }

.swagger-ui .bg-light-green, .swagger-ui .hover-bg-light-green:focus, .swagger-ui .hover-bg-light-green:hover { background-color: #228762; }

.swagger-ui .bg-navy, .swagger-ui .hover-bg-navy:focus, .swagger-ui .hover-bg-navy:hover { background-color: #0d1d35; }

.swagger-ui .bg-dark-blue, .swagger-ui .hover-bg-dark-blue:focus, .swagger-ui .hover-bg-dark-blue:hover { background-color: #20497e; }

.swagger-ui .bg-blue, .swagger-ui .hover-bg-blue:focus, .swagger-ui .hover-bg-blue:hover { background-color: #4380d0; }

.swagger-ui .bg-light-blue, .swagger-ui .hover-bg-light-blue:focus, .swagger-ui .hover-bg-light-blue:hover { background-color: #20517e; }

.swagger-ui .bg-lightest-blue, .swagger-ui .hover-bg-lightest-blue:focus, .swagger-ui .hover-bg-lightest-blue:hover { background-color: #143a52; }

.swagger-ui .bg-washed-blue, .swagger-ui .hover-bg-washed-blue:focus, .swagger-ui .hover-bg-washed-blue:hover { background-color: #0c312d; }

.swagger-ui .bg-washed-green, .swagger-ui .hover-bg-washed-green:focus, .swagger-ui .hover-bg-washed-green:hover { background-color: #0f3d2c; }

.swagger-ui .bg-washed-red, .swagger-ui .hover-bg-washed-red:focus, .swagger-ui .hover-bg-washed-red:hover { background-color: #411010; }

.swagger-ui .bg-inherit, .swagger-ui .hover-bg-inherit:focus, .swagger-ui .hover-bg-inherit:hover { background-color: inherit; }

.swagger-ui .shadow-hover { transition: all .5s cubic-bezier(.165, .84, .44, 1) 0s; }

.swagger-ui .shadow-hover::after {
    border-radius: inherit;
    box-shadow: rgba(0, 0, 0, .2) 0 0 16px 2px;
    content: "";
    height: 100%;
    left: 0;
    opacity: 0;
    position: absolute;
    top: 0;
    transition: opacity .5s cubic-bezier(.165, .84, .44, 1) 0s;
    width: 100%;
    z-index: -1;
}

.swagger-ui .bg-animate, .swagger-ui .bg-animate:focus, .swagger-ui .bg-animate:hover { transition: background-color .15s ease-in-out 0s; }

.swagger-ui .nested-links a {
    color: #99bae6;
    transition: color .15s ease-in 0s;
}

.swagger-ui .nested-links a:focus, .swagger-ui .nested-links a:hover {
    color: #a9cbea;
    transition: color .15s ease-in 0s;
}

.swagger-ui .opblock-tag {
    border-bottom: 1px solid rgba(58, 64, 80, .3);
    color: #b5bac9;
    transition: all .2s ease 0s;
}

.swagger-ui .opblock-tag svg, .swagger-ui section.models h4 svg { transition: all .4s ease 0s; }

.swagger-ui .opblock {
    border: 1px solid #000;
    border-radius: 4px;
    box-shadow: rgba(0, 0, 0, .19) 0 0 3px;
    margin: 0 0 15px;
}

.swagger-ui .opblock .tab-header .tab-item.active h4 span::after { background: gray; }

.swagger-ui .opblock.is-open .opblock-summary { border-bottom: 1px solid #000; }

.swagger-ui .opblock .opblock-section-header {
    background: rgba(28, 28, 33, .8);
    box-shadow: rgba(0, 0, 0, .1) 0 1px 2px;
}

.swagger-ui .opblock .opblock-section-header > label > span { padding: 0 10px 0 0; }

.swagger-ui .opblock .opblock-summary-method {
    background: #000;
    color: #fff;
    text-shadow: rgba(0, 0, 0, .1) 0 1px 0;
}

.swagger-ui .opblock.opblock-post {
    background: rgba(72, 203, 144, .1);
    border-color: #48cb90;
}

.swagger-ui .opblock.opblock-post .opblock-summary-method, .swagger-ui .opblock.opblock-post .tab-header .tab-item.active h4 span::after { background: #48cb90; }

.swagger-ui .opblock.opblock-post .opblock-summary { border-color: #48cb90; }

.swagger-ui .opblock.opblock-put {
    background: rgba(213, 157, 88, .1);
    border-color: #d59d58;
}

.swagger-ui .opblock.opblock-put .opblock-summary-method, .swagger-ui .opblock.opblock-put .tab-header .tab-item.active h4 span::after { background: #d59d58; }

.swagger-ui .opblock.opblock-put .opblock-summary { border-color: #d59d58; }

.swagger-ui .opblock.opblock-delete {
    background: rgba(200, 50, 50, .1);
    border-color: #c83232;
}

.swagger-ui .opblock.opblock-delete .opblock-summary-method, .swagger-ui .opblock.opblock-delete .tab-header .tab-item.active h4 span::after { background: #c83232; }

.swagger-ui .opblock.opblock-delete .opblock-summary { border-color: #c83232; }

.swagger-ui .opblock.opblock-get {
    background: rgba(42, 105, 167, .1);
    border-color: #2a69a7;
}

.swagger-ui .opblock.opblock-get .opblock-summary-method, .swagger-ui .opblock.opblock-get .tab-header .tab-item.active h4 span::after { background: #2a69a7; }

.swagger-ui .opblock.opblock-get .opblock-summary { border-color: #2a69a7; }

.swagger-ui .opblock.opblock-patch {
    background: rgba(92, 214, 188, .1);
    border-color: #5cd6bc;
}

.swagger-ui .opblock.opblock-patch .opblock-summary-method, .swagger-ui .opblock.opblock-patch .tab-header .tab-item.active h4 span::after { background: #5cd6bc; }

.swagger-ui .opblock.opblock-patch .opblock-summary { border-color: #5cd6bc; }

.swagger-ui .opblock.opblock-head {
    background: rgba(140, 63, 207, .1);
    border-color: #8c3fcf;
}

.swagger-ui .opblock.opblock-head .opblock-summary-method, .swagger-ui .opblock.opblock-head .tab-header .tab-item.active h4 span::after { background: #8c3fcf; }

.swagger-ui .opblock.opblock-head .opblock-summary { border-color: #8c3fcf; }

.swagger-ui .opblock.opblock-options {
    background: rgba(36, 89, 143, .1);
    border-color: #24598f;
}

.swagger-ui .opblock.opblock-options .opblock-summary-method, .swagger-ui .opblock.opblock-options .tab-header .tab-item.active h4 span::after { background: #24598f; }

.swagger-ui .opblock.opblock-options .opblock-summary { border-color: #24598f; }

.swagger-ui .opblock.opblock-deprecated {
    background: rgba(46, 46, 46, .1);
    border-color: #2e2e2e;
    opacity: .6;
}

.swagger-ui .opblock.opblock-deprecated .opblock-summary-method, .swagger-ui .opblock.opblock-deprecated .tab-header .tab-item.active h4 span::after { background: #2e2e2e; }

.swagger-ui .opblock.opblock-deprecated .opblock-summary { border-color: #2e2e2e; }

.swagger-ui .filter .operation-filter-input { border: 2px solid #2b3446; }

.swagger-ui .tab li:first-of-type::after { background: rgba(0, 0, 0, .2); }

.swagger-ui .download-contents {
    background: #7c8192;
    color: #fff;
}

.swagger-ui .scheme-container {
    background: #1c1c21;
    box-shadow: rgba(0, 0, 0, .15) 0 1px 2px 0;
}

.swagger-ui .loading-container .loading::before {
    animation: 1s linear 0s infinite normal none running rotation, .5s ease 0s 1 normal none running opacity;
    border-color: rgba(0, 0, 0, .6) rgba(84, 84, 84, .1) rgba(84, 84, 84, .1);
}

.swagger-ui .response-control-media-type--accept-controller select { border-color: #196619; }

.swagger-ui .response-control-media-type__accept-message { color: #99e699; }

.swagger-ui .version-pragma__message code { background-color: #3b3b3b; }

.swagger-ui .btn {
    background: 0 0;
    border: 2px solid gray;
    box-shadow: rgba(0, 0, 0, .1) 0 1px 2px;
    color: #b5bac9;
}

.swagger-ui .btn:hover { box-shadow: rgba(0, 0, 0, .3) 0 0 5px; }

.swagger-ui .btn.authorize, .swagger-ui .btn.cancel {
    background-color: transparent;
    border-color: #a72a2a;
    color: #e69999;
}

.swagger-ui .btn.authorize {
    border-color: #48cb90;
    color: #9ce3c3;
}

.swagger-ui .btn.authorize svg { fill: #9ce3c3; }

.swagger-ui .btn.execute {
    background-color: #5892d5;
    border-color: #5892d5;
    color: #fff;
}

.swagger-ui .copy-to-clipboard { background: #7c8192; }

.swagger-ui .copy-to-clipboard button { background: url("data:image/svg+xml;charset=utf-8,<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" aria-hidden=\"true\"><path fill=\"%23fff\" fill-rule=\"evenodd\" d=\"M2 13h4v1H2v-1zm5-6H2v1h5V7zm2 3V8l-3 3 3 3v-2h5v-2H9zM4.5 9H2v1h2.5V9zM2 12h2.5v-1H2v1zm9 1h1v2c-.02.28-.11.52-.3.7-.19.18-.42.28-.7.3H1c-.55 0-1-.45-1-1V4c0-.55.45-1 1-1h3c0-1.11.89-2 2-2 1.11 0 2 .89 2 2h3c.55 0 1 .45 1 1v5h-1V6H1v9h10v-2zM2 5h8c0-.55-.45-1-1-1H8c-.55 0-1-.45-1-1s-.45-1-1-1-1 .45-1 1-.45 1-1 1H3c-.55 0-1 .45-1 1z\"/></svg>") 50% center no-repeat; }

.swagger-ui select {
    background: url("data:image/svg+xml;charset=utf-8,<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\"><path d=\"M13.418 7.859a.695.695 0 01.978 0 .68.68 0 010 .969l-3.908 3.83a.697.697 0 01-.979 0l-3.908-3.83a.68.68 0 010-.969.695.695 0 01.978 0L10 11l3.418-3.141z\"/></svg>") right 10px center/20px no-repeat #212121;
    background: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjxzdmcKICAgeG1sbnM6ZGM9Imh0dHA6Ly9wdXJsLm9yZy9kYy9lbGVtZW50cy8xLjEvIgogICB4bWxuczpjYz0iaHR0cDovL2NyZWF0aXZlY29tbW9ucy5vcmcvbnMjIgogICB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiCiAgIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiAgIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIKICAgeG1sbnM6c29kaXBvZGk9Imh0dHA6Ly9zb2RpcG9kaS5zb3VyY2Vmb3JnZS5uZXQvRFREL3NvZGlwb2RpLTAuZHRkIgogICB4bWxuczppbmtzY2FwZT0iaHR0cDovL3d3dy5pbmtzY2FwZS5vcmcvbmFtZXNwYWNlcy9pbmtzY2FwZSIKICAgaW5rc2NhcGU6dmVyc2lvbj0iMS4wICg0MDM1YTRmYjQ5LCAyMDIwLTA1LTAxKSIKICAgc29kaXBvZGk6ZG9jbmFtZT0iZG93bmxvYWQuc3ZnIgogICBpZD0ic3ZnNCIKICAgdmVyc2lvbj0iMS4xIgogICB2aWV3Qm94PSIwIDAgMjAgMjAiPgogIDxtZXRhZGF0YQogICAgIGlkPSJtZXRhZGF0YTEwIj4KICAgIDxyZGY6UkRGPgogICAgICA8Y2M6V29yawogICAgICAgICByZGY6YWJvdXQ9IiI+CiAgICAgICAgPGRjOmZvcm1hdD5pbWFnZS9zdmcreG1sPC9kYzpmb3JtYXQ+CiAgICAgICAgPGRjOnR5cGUKICAgICAgICAgICByZGY6cmVzb3VyY2U9Imh0dHA6Ly9wdXJsLm9yZy9kYy9kY21pdHlwZS9TdGlsbEltYWdlIiAvPgogICAgICA8L2NjOldvcms+CiAgICA8L3JkZjpSREY+CiAgPC9tZXRhZGF0YT4KICA8ZGVmcwogICAgIGlkPSJkZWZzOCIgLz4KICA8c29kaXBvZGk6bmFtZWR2aWV3CiAgICAgaW5rc2NhcGU6Y3VycmVudC1sYXllcj0ic3ZnNCIKICAgICBpbmtzY2FwZTp3aW5kb3ctbWF4aW1pemVkPSIxIgogICAgIGlua3NjYXBlOndpbmRvdy15PSItOSIKICAgICBpbmtzY2FwZTp3aW5kb3cteD0iLTkiCiAgICAgaW5rc2NhcGU6Y3k9IjEwIgogICAgIGlua3NjYXBlOmN4PSIxMCIKICAgICBpbmtzY2FwZTp6b29tPSI0MS41IgogICAgIHNob3dncmlkPSJmYWxzZSIKICAgICBpZD0ibmFtZWR2aWV3NiIKICAgICBpbmtzY2FwZTp3aW5kb3ctaGVpZ2h0PSIxMDAxIgogICAgIGlua3NjYXBlOndpbmRvdy13aWR0aD0iMTkyMCIKICAgICBpbmtzY2FwZTpwYWdlc2hhZG93PSIyIgogICAgIGlua3NjYXBlOnBhZ2VvcGFjaXR5PSIwIgogICAgIGd1aWRldG9sZXJhbmNlPSIxMCIKICAgICBncmlkdG9sZXJhbmNlPSIxMCIKICAgICBvYmplY3R0b2xlcmFuY2U9IjEwIgogICAgIGJvcmRlcm9wYWNpdHk9IjEiCiAgICAgYm9yZGVyY29sb3I9IiM2NjY2NjYiCiAgICAgcGFnZWNvbG9yPSIjZmZmZmZmIiAvPgogIDxwYXRoCiAgICAgc3R5bGU9ImZpbGw6I2ZmZmZmZiIKICAgICBpZD0icGF0aDIiCiAgICAgZD0iTTEzLjQxOCA3Ljg1OWEuNjk1LjY5NSAwIDAxLjk3OCAwIC42OC42OCAwIDAxMCAuOTY5bC0zLjkwOCAzLjgzYS42OTcuNjk3IDAgMDEtLjk3OSAwbC0zLjkwOC0zLjgzYS42OC42OCAwIDAxMC0uOTY5LjY5NS42OTUgMCAwMS45NzggMEwxMCAxMWwzLjQxOC0zLjE0MXoiIC8+Cjwvc3ZnPgo=) right 10px center/20px no-repeat #1c1c21;
    border: 2px solid #41444e;
}

.swagger-ui select[multiple] { background: #212121; }

.swagger-ui button.invalid, .swagger-ui input[type=email].invalid, .swagger-ui input[type=file].invalid, .swagger-ui input[type=password].invalid, .swagger-ui input[type=search].invalid, .swagger-ui input[type=text].invalid, .swagger-ui select.invalid, .swagger-ui textarea.invalid {
    background: #390e0e;
    border-color: #c83232;
}

.swagger-ui input[type=email], .swagger-ui input[type=file], .swagger-ui input[type=password], .swagger-ui input[type=search], .swagger-ui input[type=text], .swagger-ui textarea {
    background: #1c1c21;
    border: 1px solid #404040;
}

.swagger-ui textarea {
    background: rgba(28, 28, 33, .8);
    color: #b5bac9;
}

.swagger-ui input[disabled], .swagger-ui select[disabled] {
    background-color: #1f1f1f;
    color: #bfbfbf;
}

.swagger-ui textarea[disabled] {
    background-color: #41444e;
    color: #fff;
}

.swagger-ui select[disabled] { border-color: #878787; }

.swagger-ui textarea:focus { border: 2px solid #2a69a7; }

.swagger-ui .checkbox input[type=checkbox] + label > .item {
    background: #303030;
    box-shadow: #303030 0 0 0 2px;
}

.swagger-ui .checkbox input[type=checkbox]:checked + label > .item { background: url("data:image/svg+xml;charset=utf-8,<svg width=\"10\" height=\"8\" viewBox=\"3 7 10 8\" xmlns=\"http://www.w3.org/2000/svg\"><path fill=\"%2341474E\" fill-rule=\"evenodd\" d=\"M6.333 15L3 11.667l1.333-1.334 2 2L11.667 7 13 8.333z\"/></svg>") 50% center no-repeat #303030; }

.swagger-ui .dialog-ux .backdrop-ux { background: rgba(0, 0, 0, .8); }

.swagger-ui .dialog-ux .modal-ux {
    background: #1c1c21;
    border: 1px solid #2e2e2e;
    box-shadow: rgba(0, 0, 0, .2) 0 10px 30px 0;
}

.swagger-ui .dialog-ux .modal-ux-header .close-modal { background: 0 0; }

.swagger-ui .model .deprecated span, .swagger-ui .model .deprecated td { color: #bfbfbf !important; }

.swagger-ui .model-toggle::after { background: url("data:image/svg+xml;charset=utf-8,<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\"><path d=\"M10 6L8.59 7.41 13.17 12l-4.58 4.59L10 18l6-6z\"/></svg>") 50% center/100% no-repeat; }

.swagger-ui .model-hint {
    background: rgba(0, 0, 0, .7);
    color: #ebebeb;
}

.swagger-ui section.models { border: 1px solid rgba(58, 64, 80, .3); }

.swagger-ui section.models.is-open h4 { border-bottom: 1px solid rgba(58, 64, 80, .3); }

.swagger-ui section.models .model-container { background: rgba(0, 0, 0, .05); }

.swagger-ui section.models .model-container:hover { background: rgba(0, 0, 0, .07); }

.swagger-ui .model-box { background: rgba(0, 0, 0, .1); }

.swagger-ui .prop-type { color: #aaaad4; }

.swagger-ui table thead tr td, .swagger-ui table thead tr th {
    border-bottom: 1px solid rgba(58, 64, 80, .2);
    color: #b5bac9;
}

.swagger-ui .parameter__name.required::after { color: rgba(230, 153, 153, .6); }

.swagger-ui .topbar .download-url-wrapper .select-label { color: #f0f0f0; }

.swagger-ui .topbar .download-url-wrapper .download-url-button {
    background: #63a040;
    color: #fff;
}

.swagger-ui .info .title small { background: #7c8492; }

.swagger-ui .info .title small.version-stamp { background-color: #7a9b27; }

.swagger-ui .auth-container .errors {
    background-color: #350d0d;
    color: #b5bac9;
}

.swagger-ui .errors-wrapper {
    background: rgba(200, 50, 50, .1);
    border: 2px solid #c83232;
}

.swagger-ui .markdown code, .swagger-ui .renderedmarkdown code {
    background: rgba(0, 0, 0, .05);
    color: #c299e6;
}

.swagger-ui .model-toggle:after { background: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjxzdmcKICAgeG1sbnM6ZGM9Imh0dHA6Ly9wdXJsLm9yZy9kYy9lbGVtZW50cy8xLjEvIgogICB4bWxuczpjYz0iaHR0cDovL2NyZWF0aXZlY29tbW9ucy5vcmcvbnMjIgogICB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiCiAgIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiAgIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIKICAgeG1sbnM6c29kaXBvZGk9Imh0dHA6Ly9zb2RpcG9kaS5zb3VyY2Vmb3JnZS5uZXQvRFREL3NvZGlwb2RpLTAuZHRkIgogICB4bWxuczppbmtzY2FwZT0iaHR0cDovL3d3dy5pbmtzY2FwZS5vcmcvbmFtZXNwYWNlcy9pbmtzY2FwZSIKICAgaW5rc2NhcGU6dmVyc2lvbj0iMS4wICg0MDM1YTRmYjQ5LCAyMDIwLTA1LTAxKSIKICAgc29kaXBvZGk6ZG9jbmFtZT0iZG93bmxvYWQyLnN2ZyIKICAgaWQ9InN2ZzQiCiAgIHZlcnNpb249IjEuMSIKICAgaGVpZ2h0PSIyNCIKICAgd2lkdGg9IjI0Ij4KICA8bWV0YWRhdGEKICAgICBpZD0ibWV0YWRhdGExMCI+CiAgICA8cmRmOlJERj4KICAgICAgPGNjOldvcmsKICAgICAgICAgcmRmOmFib3V0PSIiPgogICAgICAgIDxkYzpmb3JtYXQ+aW1hZ2Uvc3ZnK3htbDwvZGM6Zm9ybWF0PgogICAgICAgIDxkYzp0eXBlCiAgICAgICAgICAgcmRmOnJlc291cmNlPSJodHRwOi8vcHVybC5vcmcvZGMvZGNtaXR5cGUvU3RpbGxJbWFnZSIgLz4KICAgICAgPC9jYzpXb3JrPgogICAgPC9yZGY6UkRGPgogIDwvbWV0YWRhdGE+CiAgPGRlZnMKICAgICBpZD0iZGVmczgiIC8+CiAgPHNvZGlwb2RpOm5hbWVkdmlldwogICAgIGlua3NjYXBlOmN1cnJlbnQtbGF5ZXI9InN2ZzQiCiAgICAgaW5rc2NhcGU6d2luZG93LW1heGltaXplZD0iMSIKICAgICBpbmtzY2FwZTp3aW5kb3cteT0iLTkiCiAgICAgaW5rc2NhcGU6d2luZG93LXg9Ii05IgogICAgIGlua3NjYXBlOmN5PSIxMiIKICAgICBpbmtzY2FwZTpjeD0iMTIiCiAgICAgaW5rc2NhcGU6em9vbT0iMzQuNTgzMzMzIgogICAgIHNob3dncmlkPSJmYWxzZSIKICAgICBpZD0ibmFtZWR2aWV3NiIKICAgICBpbmtzY2FwZTp3aW5kb3ctaGVpZ2h0PSIxMDAxIgogICAgIGlua3NjYXBlOndpbmRvdy13aWR0aD0iMTkyMCIKICAgICBpbmtzY2FwZTpwYWdlc2hhZG93PSIyIgogICAgIGlua3NjYXBlOnBhZ2VvcGFjaXR5PSIwIgogICAgIGd1aWRldG9sZXJhbmNlPSIxMCIKICAgICBncmlkdG9sZXJhbmNlPSIxMCIKICAgICBvYmplY3R0b2xlcmFuY2U9IjEwIgogICAgIGJvcmRlcm9wYWNpdHk9IjEiCiAgICAgYm9yZGVyY29sb3I9IiM2NjY2NjYiCiAgICAgcGFnZWNvbG9yPSIjZmZmZmZmIiAvPgogIDxwYXRoCiAgICAgc3R5bGU9ImZpbGw6I2ZmZmZmZiIKICAgICBpZD0icGF0aDIiCiAgICAgZD0iTTEwIDZMOC41OSA3LjQxIDEzLjE3IDEybC00LjU4IDQuNTlMMTAgMThsNi02eiIgLz4KPC9zdmc+Cg==) 50% no-repeat; }

.swagger-ui .expand-operation svg, .swagger-ui section.models h4 svg { fill: #fff; }

::-webkit-scrollbar-track { background-color: #646464 !important; }

::-webkit-scrollbar-thumb {
    background-color: #242424 !important;
    border: 2px solid #3e4346 !important;
}

::-webkit-scrollbar-button:vertical:start:decrement {
    background: linear-gradient(130deg, #696969 40%, rgba(255, 0, 0, 0) 41%), linear-gradient(230deg, #696969 40%, transparent 41%), linear-gradient(0deg, #696969 40%, transparent 31%);
    background-color: #b6b6b6;
}

::-webkit-scrollbar-button:vertical:end:increment {
    background: linear-gradient(310deg, #696969 40%, transparent 41%), linear-gradient(50deg, #696969 40%, transparent 41%), linear-gradient(180deg, #696969 40%, transparent 31%);
    background-color: #b6b6b6;
}

::-webkit-scrollbar-button:horizontal:end:increment {
    background: linear-gradient(210deg, #696969 40%, transparent 41%), linear-gradient(330deg, #696969 40%, transparent 41%), linear-gradient(90deg, #696969 30%, transparent 31%);
    background-color: #b6b6b6;
}

::-webkit-scrollbar-button:horizontal:start:decrement {
    background: linear-gradient(30deg, #696969 40%, transparent 41%), linear-gradient(150deg, #696969 40%, transparent 41%), linear-gradient(270deg, #696969 30%, transparent 31%);
    background-color: #b6b6b6;
}

::-webkit-scrollbar-button, ::-webkit-scrollbar-track-piece { background-color: #3e4346 !important; }

.swagger-ui .black, .swagger-ui .checkbox, .swagger-ui .dark-gray, .swagger-ui .download-url-wrapper .loading, .swagger-ui .errors-wrapper .errors small, .swagger-ui .fallback, .swagger-ui .filter .loading, .swagger-ui .gray, .swagger-ui .hover-black:focus, .swagger-ui .hover-black:hover, .swagger-ui .hover-dark-gray:focus, .swagger-ui .hover-dark-gray:hover, .swagger-ui .hover-gray:focus, .swagger-ui .hover-gray:hover, .swagger-ui .hover-light-silver:focus, .swagger-ui .hover-light-silver:hover, .swagger-ui .hover-mid-gray:focus, .swagger-ui .hover-mid-gray:hover, .swagger-ui .hover-near-black:focus, .swagger-ui .hover-near-black:hover, .swagger-ui .hover-silver:focus, .swagger-ui .hover-silver:hover, .swagger-ui .light-silver, .swagger-ui .markdown pre, .swagger-ui .mid-gray, .swagger-ui .model .property, .swagger-ui .model .property.primitive, .swagger-ui .model-title, .swagger-ui .near-black, .swagger-ui .parameter__extension, .swagger-ui .parameter__in, .swagger-ui .prop-format, .swagger-ui .renderedmarkdown pre, .swagger-ui .response-col_links .response-undocumented, .swagger-ui .response-col_status .response-undocumented, .swagger-ui .silver, .swagger-ui section.models h4, .swagger-ui section.models h5, .swagger-ui span.token-not-formatted, .swagger-ui span.token-string, .swagger-ui table.headers .header-example, .swagger-ui table.model tr.description, .swagger-ui table.model tr.extension { color: #bfbfbf; }

.swagger-ui .hover-white:focus, .swagger-ui .hover-white:hover, .swagger-ui .info .title small pre, .swagger-ui .topbar a, .swagger-ui .white { color: #fff; }

.swagger-ui .bg-black-10, .swagger-ui .hover-bg-black-10:focus, .swagger-ui .hover-bg-black-10:hover, .swagger-ui .stripe-dark:nth-child(2n + 1) { background-color: rgba(0, 0, 0, .1); }

.swagger-ui .bg-white-10, .swagger-ui .hover-bg-white-10:focus, .swagger-ui .hover-bg-white-10:hover, .swagger-ui .stripe-light:nth-child(2n + 1) { background-color: rgba(28, 28, 33, .1); }

.swagger-ui .bg-light-silver, .swagger-ui .hover-bg-light-silver:focus, .swagger-ui .hover-bg-light-silver:hover, .swagger-ui .striped--light-silver:nth-child(2n + 1) { background-color: #6e6e6e; }

.swagger-ui .bg-moon-gray, .swagger-ui .hover-bg-moon-gray:focus, .swagger-ui .hover-bg-moon-gray:hover, .swagger-ui .striped--moon-gray:nth-child(2n + 1) { background-color: #4d4d4d; }

.swagger-ui .bg-light-gray, .swagger-ui .hover-bg-light-gray:focus, .swagger-ui .hover-bg-light-gray:hover, .swagger-ui .striped--light-gray:nth-child(2n + 1) { background-color: #2b2b2b; }

.swagger-ui .bg-near-white, .swagger-ui .hover-bg-near-white:focus, .swagger-ui .hover-bg-near-white:hover, .swagger-ui .striped--near-white:nth-child(2n + 1) { background-color: #242424; }

.swagger-ui .opblock-tag:hover, .swagger-ui section.models h4:hover { background: rgba(0, 0, 0, .02); }

.swagger-ui .checkbox p, .swagger-ui .dialog-ux .modal-ux-content h4, .swagger-ui .dialog-ux .modal-ux-content p, .swagger-ui .dialog-ux .modal-ux-header h3, .swagger-ui .errors-wrapper .errors h4, .swagger-ui .errors-wrapper hgroup h4, .swagger-ui .info .base-url, .swagger-ui .info .title, .swagger-ui .info h1, .swagger-ui .info h2, .swagger-ui .info h3, .swagger-ui .info h4, .swagger-ui .info h5, .swagger-ui .info li, .swagger-ui .info p, .swagger-ui .info table, .swagger-ui .loading-container .loading::after, .swagger-ui .model, .swagger-ui .opblock .opblock-section-header h4, .swagger-ui .opblock .opblock-section-header > label, .swagger-ui .opblock .opblock-summary-description, .swagger-ui .opblock .opblock-summary-operation-id, .swagger-ui .opblock .opblock-summary-path, .swagger-ui .opblock .opblock-summary-path__deprecated, .swagger-ui .opblock-description-wrapper, .swagger-ui .opblock-description-wrapper h4, .swagger-ui .opblock-description-wrapper p, .swagger-ui .opblock-external-docs-wrapper, .swagger-ui .opblock-external-docs-wrapper h4, .swagger-ui .opblock-external-docs-wrapper p, .swagger-ui .opblock-tag small, .swagger-ui .opblock-title_normal, .swagger-ui .opblock-title_normal h4, .swagger-ui .opblock-title_normal p, .swagger-ui .parameter__name, .swagger-ui .parameter__type, .swagger-ui .response-col_links, .swagger-ui .response-col_status, .swagger-ui .responses-inner h4, .swagger-ui .responses-inner h5, .swagger-ui .scheme-container .schemes > label, .swagger-ui .scopes h2, .swagger-ui .servers > label, .swagger-ui .tab li, .swagger-ui label, .swagger-ui select, .swagger-ui table.headers td { color: #b5bac9; }

.swagger-ui .download-url-wrapper .failed, .swagger-ui .filter .failed, .swagger-ui .model-deprecated-warning, .swagger-ui .parameter__deprecated, .swagger-ui .parameter__name.required span, .swagger-ui table.model tr.property-row .star { color: #e69999; }

.swagger-ui .opblock-body pre.microlight, .swagger-ui textarea.curl {
    background: #41444e;
    border-radius: 4px;
    color: #fff;
}

.swagger-ui .expand-methods svg, .swagger-ui .expand-methods:hover svg { fill: #bfbfbf; }

.swagger-ui .auth-container, .swagger-ui .dialog-ux .modal-ux-header { border-bottom: 1px solid #2e2e2e; }

.swagger-ui .topbar .download-url-wrapper .select-label select, .swagger-ui .topbar .download-url-wrapper input[type=text] { border: 2px solid #63a040; }

.swagger-ui .info a, .swagger-ui .info a:hover, .swagger-ui .scopes h2 a { color: #99bde6; }

/* Dark Scrollbar */
::-webkit-scrollbar {
    width: 14px;
    height: 14px;
}

::-webkit-scrollbar-button {
    background-color: #3e4346 !important;
}

::-webkit-scrollbar-track {
    background-color: #646464 !important;
}

::-webkit-scrollbar-track-piece {
    background-color: #3e4346 !important;
}

::-webkit-scrollbar-thumb {
    height: 50px;
    background-color: #242424 !important;
    border: 2px solid #3e4346 !important;
}

::-webkit-scrollbar-corner {}

::-webkit-resizer {}

::-webkit-scrollbar-button:vertical:start:decrement {
    background:
        linear-gradient(130deg, #696969 40%, rgba(255, 0, 0, 0) 41%),
        linear-gradient(230deg, #696969 40%, rgba(0, 0, 0, 0) 41%),
        linear-gradient(0deg, #696969 40%, rgba(0, 0, 0, 0) 31%);
    background-color: #b6b6b6;
}

::-webkit-scrollbar-button:vertical:end:increment {
    background:
        linear-gradient(310deg, #696969 40%, rgba(0, 0, 0, 0) 41%),
        linear-gradient(50deg, #696969 40%, rgba(0, 0, 0, 0) 41%),
        linear-gradient(180deg, #696969 40%, rgba(0, 0, 0, 0) 31%);
    background-color: #b6b6b6;
}

::-webkit-scrollbar-button:horizontal:end:increment {
    background:
        linear-gradient(210deg, #696969 40%, rgba(0, 0, 0, 0) 41%),
        linear-gradient(330deg, #696969 40%, rgba(0, 0, 0, 0) 41%),
        linear-gradient(90deg, #696969 30%, rgba(0, 0, 0, 0) 31%);
    background-color: #b6b6b6;
}

::-webkit-scrollbar-button:horizontal:start:decrement {
    background:
        linear-gradient(30deg, #696969 40%, rgba(0, 0, 0, 0) 41%),
        linear-gradient(150deg, #696969 40%, rgba(0, 0, 0, 0) 41%),
        linear-gradient(270deg, #696969 30%, rgba(0, 0, 0, 0) 31%);
    background-color: #b6b6b6;
}
	`
