class App extends React.Component {
    render() {
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
}

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Jokeish</h1>
                    <p>A load of Dad jokes XD</p>
                    <p>Sign in to get access</p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-login btn-block">Sign in</a>
                </div>
            </div>
        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            jokes: []
        }
    }

    render() {
        return(
            <div className="container">
                <div className="col-lg-12">
                    <br />
                    <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                    <h2>Jokeish</h2>
                    <p>Lets feed you with some funny jokes</p>
                    <div className="row">
                        {this.state.jokes.map(function(joke, i) {
                            return (<Joke key={i} joke={joke} />);
                        })}
                    </div>
                </div>
            </div>
        )
    }
}

class Joke extends React.Component {

}