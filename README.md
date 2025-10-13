Power4 Web — Jeu Puissance 4 en Go

Bienvenue dans le projet Power4 Web, qui est une application web développée en Go (Golang) qui permet à deux joueurs de s’affronter dans une partie de Puissance 4 sur le même ordinateur, grâce à une interface web simple.

Ce projet est conçu pour être facile à utiliser, même si vous n’avez jamais codé. Chaque étape est expliquée pour que vous puissiez le lancer sans difficultée.

    🧠 Objectif du projet:
        Créer un serveur web en Go qui va:

        Afficher une interface HTML pour jouer à Puissance 4.
        Gérer toute la logique du jeu côté serveur.
        Permettre de jouer à deux en local.
        Proposer plusieurs niveaux de difficulté, de facile a difficile, on aura un changement de taille de la grille.
        Ajouter la gravité inversée au bout de 5 tours.

    📦 Les prérequis:
        Avant de commencer, vous devez avoir :
        Go installé sur votre ordinateur : Télécharger Go
        Un navigateur web (par exemple: Chrome, Firefox, etc.)
        Un terminal (comme: cmd, PowerShell, Terminal Mac/Linux)

    📁 Structure du projet:
        Voici les dossiers et fichiers que vous trouverez :

            Power4-Web/
            ├── models/           # Contient la logique du jeu
            ├── static/           # Contient le fichier CSS pour le design
            │   └── style.css
            ├── templates/        # Contient les pages HTML
            │   ├── start.html    # Page d’accueil
            │   ├── game.html     # Page du jeu
            │   ├── win.html      # Page de victoire
            │   ├── draw.html     # Page de match nul
            │   └── history.html  # Page d'historique des parties
            ├── server.go         # Fichier principal du serveur (celui à lancer)
            ├── go.mod            # Fichier de configuration Go
            └── README.md         # Le fichier d’explication
    
    🚀 Étapes pour lancer le projet:

        1. Cloner le projet
        Ouvrez votre terminal et tapez :

        git clone https://github.com/ludivine25/Power-4.git
        cd Power-4/

        2. Initialiser le module Go
        Tapez :

        go mod init power4
        (Cela crée un fichier go.mod qui permet à Go de gérer les dépendances.)

        3. Lancer le serveur
        Tapez :

        go run server.go
        (Le serveur démarre et affiche un message comme :

        Serveur lancé sur http://localhost:8000 )

        4. Ouvrir le jeu dans le navigateur
        Copiez l’adresse affichée (http://localhost:8000) et collez-la dans votre navigateur.

    🕹️ Comment jouer:

        La page d’accueil: vous devez entrer les noms des deux joueurs et choisir une difficulté :

            Facile : grille 6x7
            Normal : grille 6x9
            Difficile : grille 7x8

        Pour la page de jeu, vous devez cliquer sur une colonne pour y placer un jeton. 
        Le tour va passer automatiquement à l’autre joueur.
        Pour ce qui cooncerna le victoire ou le match nul, le jeu va le détecter automatiquement :

            Une victoire c'est quand 4 jetons sont alignés et pour un match nul c'est quand la grille est pleine, donc pas de gagnant.

        Le rematch il faut cliquer sur “Rejouez avec les mêmes paramètres” pour relancer une partie avec le même joueur.
        L'historique vous permez de consulter les parties précédentes avec les scores et qui a gagné.

    🧲 Bonus : Gravité inversée
        Tous les 5 tours, la gravité change :

        Au debut les jetons tombent du bas vers le haut, puis au bout de 5 tours le design change pour indiquer cette inversion.

    🌐 Les routes techniques (pour les développeurs):
        Méthode	    Route	    Description
        GET	          /	        Page d’accueil
        POST	     /start	    Lance une nouvelle partie
        POST	    /play	    Enregistre un coup
        GET	        /win	    Affiche la victoire
        GET	        /draw	    Affiche le match nul
        GET	        /rematch	Relance une partie
        GET	        /history	Affiche l’historique
        GET	        /static/	Sert les fichiers CSS

    🧱 Les technologies utilisées:
        Technologie	    Rôle
        Go	            Serveur web, logique du jeu, génération dynamique des pages
        HTML	        Structure des pages
        CSS	            Design visuel           

    📚 Ressources pour apprendre: 
        Documentation:
            Go net/http
            Go Templates
            HTML/CSS pour débutants

Auteur
    Ludivine G. 
    Ce projet a été réalisé dans le cadre d’un exercice de développement web en Go.