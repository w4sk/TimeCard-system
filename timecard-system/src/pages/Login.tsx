import '../Login.css'

const Login = () => {
  return (
    <main className='form-signin'>
      <form>
        <h1 className='h3 mb-3 fw-normal'>Please Sign In</h1>
        <div className='form-floating'>
          <input type='email' className='form-control' id='floating-input' placeholder='name@example.com' />
          <label htmlFor='floatingInput'>Email Adress</label>
        </div>
        <div className='form-floating'>
          <input type='password' className='form-control' id='floatingPassword' placeholder='Password'/>
          <label htmlFor='floatingPassword'>Password</label>
        </div>

        <button className='w-100 btn btn-lg btn-primary' type='submit'>Sign in</button>
      </form>
    </main>
  )
}

export default Login