#include <linux/module.h>
// Para usar KERN_INFO
#include <linux/kernel.h>

//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
// for copy_from_user
#include <asm/uaccess.h>
//Header para usar la lib seq_file y manejar el archivo en /proc
#include <linux/seq_file.h>



static int escribir_archivo(struct seq_file *archivo, void *v){
    seq_printf(archivo, "PRUEBA");
    return 0; 
}

static int al_abrir(struct inode *inode, struct file *file){
    return single_open(file, escribir_archivo, NULL);
}

// Si el kernel no es 5.6 o mayor dara errores
static struct proc_ops operaciones = {
    .proc_open = al_abrir,
    .proc_read = seq_read
};

static int _insert(void){
    proc_create("cpu_201902663",0,NULL,&operaciones);
    printk(KERN_INFO "Mensaje al insertar el modulo, CPU 201902663\n");

    return 0;
}

static void _remove(void){
    remove_proc_entry("cpu_201902663",NULL);
    printk(KERN_INFO "Mensaje al remover el modulo, CPU 201902663\n");
}

module_init(_insert);
module_exit(_remove);


// sudo dmesg -C       = limpia los mensajes