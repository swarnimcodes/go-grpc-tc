let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/github/go/go-grpc-tc
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +1 cmd/server/server.go
badd +19 cmd/client/client.go
badd +6 term://~/github/go/go-grpc-tc//73739:/bin/bash
badd +1 term://~/github/go/go-grpc-tc//73753:/bin/bash
badd +1 server_test.go
argglobal
%argdel
edit cmd/server/server.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
wincmd _ | wincmd |
split
1wincmd k
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 103 + 86) / 172)
exe '2resize ' . ((&lines * 17 + 18) / 37)
exe 'vert 2resize ' . ((&columns * 68 + 86) / 172)
exe '3resize ' . ((&lines * 17 + 18) / 37)
exe 'vert 3resize ' . ((&columns * 68 + 86) / 172)
argglobal
balt server_test.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 15 - ((14 * winheight(0) + 17) / 35)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 15
normal! 035|
wincmd w
argglobal
if bufexists(fnamemodify("term://~/github/go/go-grpc-tc//73739:/bin/bash", ":p")) | buffer term://~/github/go/go-grpc-tc//73739:/bin/bash | else | edit term://~/github/go/go-grpc-tc//73739:/bin/bash | endif
if &buftype ==# 'terminal'
  silent file term://~/github/go/go-grpc-tc//73739:/bin/bash
endif
balt cmd/client/client.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 3 - ((2 * winheight(0) + 8) / 17)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 3
normal! 0
wincmd w
argglobal
if bufexists(fnamemodify("term://~/github/go/go-grpc-tc//73753:/bin/bash", ":p")) | buffer term://~/github/go/go-grpc-tc//73753:/bin/bash | else | edit term://~/github/go/go-grpc-tc//73753:/bin/bash | endif
if &buftype ==# 'terminal'
  silent file term://~/github/go/go-grpc-tc//73753:/bin/bash
endif
balt term://~/github/go/go-grpc-tc//73739:/bin/bash
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 3 - ((2 * winheight(0) + 8) / 17)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 3
normal! 0
wincmd w
exe 'vert 1resize ' . ((&columns * 103 + 86) / 172)
exe '2resize ' . ((&lines * 17 + 18) / 37)
exe 'vert 2resize ' . ((&columns * 68 + 86) / 172)
exe '3resize ' . ((&lines * 17 + 18) / 37)
exe 'vert 3resize ' . ((&columns * 68 + 86) / 172)
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
