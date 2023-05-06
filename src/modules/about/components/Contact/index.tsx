import {TextInput} from './TextInput'
import { TextArea } from "./TextArea";
import cx from 'classnames';

export default function Contact() {
  return (
    <form className="mb-10">
      <div className="mx-auto flex flex-col justify-center gap-3">
        <div className="flex flex-col md:flex-row gap-3 items-center w-full">
          <TextInput label="Name" type="text"/>
          <TextInput label="Email" type="email"/>
        </div>

        <TextArea label='Message' className="min-h-[5rem]" placeholder="" />

        <button 
          className={cx(
            "bg-brand-100 text-brand-200 uppercase", 
            "mt-2 w-max mx-auto",
            "px-4 py-2 rounded-md",
            "hover:brightness-110 transition-all"
          )}>
          Send
        </button>
      </div>
    </form>
  )

}

