import React, { useState } from 'react'
import { useForm } from '@mantine/form';
import { Button, Group, Modal } from '@mantine/core';

const AddTodo = () => {

    const [open, setOpen] = useState(false);

    const form = useForm({
        initialValue: {
            title: "",
            body: ""
        }
    });

    return (
        <>
            <Modal
                opened={open}
                onClose={() => setOpen(false)}>
            </Modal>

            <Group
                
            >
                <Button fullWidth mb={12} onClick={() => setOpen(true)}>
                    Add todo
                </Button>
            </Group>
        </>

    )
}

export default AddTodo