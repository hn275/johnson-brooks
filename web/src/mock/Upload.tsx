import { FormEvent, useState } from "react";

export function UpLoad() {
  const [files, setFile] = useState<FileList | null>(null);

  async function handleSubmit(e: FormEvent<HTMLElement>) {
    e.preventDefault();
    if (!files) return;
    const file = files[0]!;

    try {
      const data = await processImage(file);
      console.log(data);
      // TODO: set this to backend
    } catch (e) {
      console.log(e);
    } finally {
      // loading state
    }
  }
  return (
    <form onSubmit={handleSubmit}>
      <input
        type="file"
        id="thing"
        onChange={(e) => setFile(() => e.target.files)}
      />
      <button type="submit">submit</button>
    </form>
  );
}

function processImage(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onloadend = () => {
      const imageData = (reader.result as string).split(";")[1];
      if (!imageData) return reject("invalid data");
      const base64 = imageData.split(",")[1]!;
      resolve(base64);
    };
    reader.onerror = () => reject("can't read file");
    reader.readAsDataURL(file);
  });
}
