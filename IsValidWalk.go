
package main


func IsValidWalk(walks []rune) bool {
 Ncount := 0
 Scount := 0
  Ecount := 0
  Wcount := 0
  if len(walks) != 10{
    return false
  }
  for _, walk := range walks{
    if walk == 'n'{
      Ncount ++
    }else if walk == 's' {
      Scount ++
    }else if walk == 'e' {
      Ecount ++
    }else if walk == 'w' {
      Wcount ++
    }
    
  }
  if Ncount == Scount && Wcount == Ecount {
    return true
  } else {
    return false
  }
}
