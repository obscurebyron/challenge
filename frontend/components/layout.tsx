import Alert from './alert'
import Footer from './footer'
import Meta from './meta'

type Props = {
  isAuthenticated?: boolean
  children: React.ReactNode
}

const Layout = ({ isAuthenticated, children }: Props) => {
  return (
    <>
      <Meta />
      <div className="min-h-screen">
        <Alert isAuthenticated={isAuthenticated} />
        <main>{children}</main>
      </div>
      <Footer />
    </>
  )
}

export default Layout
