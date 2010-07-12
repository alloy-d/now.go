# Be productive.

To-do lists are supposed to help you get things done.  And I suppose
looking through all the stuff you still have to do each time you want to
decide what to do next *might* be helpful.

More likely, though, it's just a time-consuming task that serves to
stress and demotivate you while sending your brain off thinking about a
*n* different things at once.  That's how it works for me, anyway.

## "Hey, me too!"

Maybe you've heard the suggestion that you paginate your to-do
list--break off, say, ten things, and just look at that sub-list until
it's done.  Now you have a smaller to-do list.  Hoorah.

Here's a better idea: divide your to-do list into two categories.  One,
all the thing (yes, thing) you should be working on now.  Two, all the
things you *aren't* working on now, and don't need to
think/worry/wonder/*whatever* about until you are.

# `include "superpowers"`

This here `now.go` gives you a little bit of magic.  Here's how it
works:

## Start it up.

    $ nowserve
    2010/07/08 12:34:17 web.go serving :5939

## When you have something you need to do later, tell it.

    $ later make sure now.go actually works.
    $ later write about now.go.

## When you want to know what you should be doing, ask it.

    $ now
    make sure now.go actually works.

## Let it know when you're done.

    $ now -done
    write about now.go.

# Just add focus and consistency.

Trust it to keep track of everything for `later`, and you'll have all
your attention for `now`.  It's (theoretically) that easy.

# No command line required!

(Note: that heading is a lie!  The webpage is still half
useless, since `later` is unaccounted for.)

GUI type, eh?  Well, then, point your browser to `localhost:5939` for a
no-nonsense, this-is-what-you-should-be-doing-now interface.
(Theoretically, a great default new-tab page!)

# Speaks HTTP, too!

Are you a robot who's just not very good at multitasking?

    POST /later/    thing=some+thing+to+do
    GET  /now.json
    POST /now       thing=go+directly+to+now
    POST /done      done=

or

    DELETE /now

(The format is obviously a little goofy and subject to change.  The
spirit will remain the same, though.)

# "I don't trust it."/"I want to reorder."/"I miss that overwhelming list-full-of-stuff feeling."

With a slight admonition that "you're probably defeating the intended
purpose", I point you to `~/.next.things`.  Verify/reorder/gaze upon to
your heart's content.  It's just plain text, one thing to a line.
Whatever changes you make will be accomodated by `now.go` (just be sure
to leave it alone while you're doing your manual editing).

# "Okay, I want it!"

Great!

First, you'll need an installation of the [Go](http://golang.org)
language.  If you don't have one already, it's easy enough to [get
one](http://golang.org/doc/install.html).

Once you've got Go up and running, you can `goinstall` now.go:

    $ goinstall github.com/alloy-d/now.go

...but this just gives you the library that the tools use.  To get the
tools themselves installed, there are a few more steps:

    $ cd $GOROOT/src/pkg/github.com/alloy-d/now.go/tools/
    $ for i in *; do (cd $i; make install); done

This will install `now`, `later`, and `servenow` to your `$GOBIN`,
which should already be in your `$PATH`.

## A special note on installation

The given method will leave the source tree intact in
`$GOROOT/src/pkg/github.com/alloy-d/now.go`.  If you don't use
`goinstall`, you will want to make sure to keep the source around after
you build the library.  At the moment, it uses an ugly hack to find the
static files needed for the HTML interface (see
[`dirty-rotten-hacks.go`](now.go/blob/master/dirty-rotten-hacks.go)).

# Miscellaneous topics

## "What is this HTTP nonsense?  How brain-dead are you?"

If you'd like something less bloated than an HTTP server, you can
replace local `now` and `later` with these lines in your
`~/.whatever-shellrc`:

    alias now='head -n 1 ~/.next.things'

    later() {
        echo "$@" >> ~/.next.things
    }

    nowdone() {
        things=$(wc -l ~/.next.things)
        tail -n $(($things-1)) ~/.next.things > ~/.next.things
    }

Now you can go along your merry way, blissfully unaware of my bloat,
bugs, poor design, and any other harmful things contained herein.

## "'Grats on reinventing the FIFO queue."

I know, I'm pretty much a champion.

Please don't nominate me for the Turing Award now, though; I have hopes
that I might yet do something *even better*.

## "Installation sucks."

I agree.  `goinstall` is an experiment in *package* installation,
though, and doesn't seem to consider packages that come with commands.
No good solution is immediately obvious to me.
