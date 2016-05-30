-module(tut1).
-export([add/2]).
-export([convert/2,convert/1]).
-export([list_length/1]).
%求最大值
-export([list_max/1]).
%翻转列表
-export([reverse/1]).

add(Aa,Ba) ->
    Aa + Ba.

convert(M,inch)->
    M/2.54;

convert(N,centimeter)->
    N*2.54.

convert({inch,Y})->
    {centimeter,Y*2.54}.

list_length([])->
    0;
list_length([FIRST|REST]) ->
    list_length(REST)+1.

list_max([Head|Rest])->
    list_max(Rest,Head).

list_max([],Max)->
    Max;
list_max([Head|Rest],Max) when Head>Max ->
    list_max(Rest,Head);
list_max([Head|Rest],Max) ->
    list_max(Rest,Max).

reverse(list)->
    reverse(list,[]).

reverse([Head|Rest],Reverse_list)->
    reverse(rest,[Head|Reverse_list]);
reverse([],Reverse_list) ->
    Reverse_list.



    

	       
