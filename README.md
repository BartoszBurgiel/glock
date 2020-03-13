# Glock
Glock is a cool tool for displaying clocks written in Golang

## Digital clock
Glock can display a digital clock showing the current time, today's date as well as the year's day in various bases including:

* base16 -> hexadecimal
* base10 -> decimal 
* base8  -> octal 
* base3  -> trinary
* base2  -> binary

It makes a quite cool widget

### Preview
<pre><code>
8 |26:30:55
3 |211:220:1200
2 |10110:11000:101101

0d/03/7e4
49th day of the year
</code></pre>

## Analog clock 
Glock can also display some data (similarly to a pie chart) or just the time as a digital clock: 

### Preview 
<pre><code>
 -------   -------   ------- 
| \     | |       | |       |
|   .   | |   .   | |   .   |
|       | |     \ | |     \ |
 -------   -------   ------- 
  hours  | minutes | seconds 
</code></pre>