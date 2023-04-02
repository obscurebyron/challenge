import Container from './container'
import cn from 'classnames'
import Link from 'next/link'

type Props = {
  isAuthenticated?: boolean
}

const Alert = ({ isAuthenticated }: Props) => {
  return (
    <div
      className={cn('border-b', {
        'bg-neutral-800 border-neutral-800 text-white': isAuthenticated,
        'bg-neutral-50 border-neutral-200': !isAuthenticated,
      })}
    >
      <Container>
        <div className="py-2 text-center text-sm">
          
            {
              isAuthenticated ?
                <div>
                  You are currently authenticated {' '}
                  <Link href="/login" className="underline hover:text-blue-600 duration-200 transition-colors">Logout</Link>
                </div>
              :
                <div>
                  You are currently unauthenticated. {' '}
                  <Link href="/login" className="underline hover:text-blue-600 duration-200 transition-colors">Authenticate</Link>
                </div>
            }

        </div>
      </Container>
    </div>
  )
}

export default Alert
