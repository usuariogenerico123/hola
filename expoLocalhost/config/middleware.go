package config

import (
	"fmt"
	"net/http"
	"strings"
)




func Filter(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		
		if(strings.Contains(path, "/.")){
			w.WriteHeader(http.StatusNoContent)
			
			fmt.Fprintf(w, "UNAUTHORIZED")
			return
		}
		privatePath := strings.Split(path, "")
		if(len(privatePath)>1){
			if(privatePath[1] == "."){ 
				//Aca se puede usar config.WhiteFiles para filtrar que archivos deben estar protegidos
				//Por el momento todos .env .git .vscode etc etc estan protegidos
				fmt.Fprintf(w, "UNAUTHORIZED")
				return
			}
		}
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}

