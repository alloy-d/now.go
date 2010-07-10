# Be productive.

To-do lists are supposed to help you get things done.  And I suppose
looking through all the stuff you still have to do each time you want to
decide what to do next *might* be helpful.

More likely, though, it's just a time-consuming task that serves to
stress and demotivate you while sending your brain off thinking about a
*n* different things at once.  That's how it works for me, anyway.

## "Hey, me too!"

One suggestion is to paginate your to-do list--break off, say, ten
things, and just look at that sub-list until it's done.  Now you have a
smaller to-do list.  Hoorah.

Here's a better idea: divide your to-do list into two categories.  One,
all the thing (yes, thing) you should be working on now.  Two, all the
things you *aren't* working on now, and don't need to
think/worry/wonder/*whatever* about until you are.

# `include "superpowers"`

This here `now.go` gives you a little bit of magic.  Here's how it
works:

## Start it up.

    $ serve
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

# Miscellaneous topics

## "What is this HTTP nonsense?  How brain-dead are you?"

If you'd like something less bloated than an HTTP server, you can
replace local `now` and `later` with these lines in your
`~/.whatever-shellrc`:

    alias now='head -n 1 ~/.next.things'

    later() {
        echo "$@" >> ~/.next.things
    }

Now you can go along your merry way, blissfully unaware of my bloat,
bugs, poor design, and any other harmful things contained herein.

## "'Grats on reinventing the FIFO queue."

I know, I'm pretty much a champion.

Please don't nominate me for the Turing Award now, though; I have hopes
that I might yet do something *even better*.
